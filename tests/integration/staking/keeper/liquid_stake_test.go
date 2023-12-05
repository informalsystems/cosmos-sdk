package keeper_test

import (
	"testing"
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	accountkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/testutil"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

// Helper function to clear the Bonded pool balances before a unit test
func clearPoolBalance(t *testing.T, sk keeper.Keeper, ak accountkeeper.AccountKeeper, bk bankkeeper.Keeper, ctx sdk.Context) {
	bondDenom := sk.BondDenom(ctx)
	initialBondedBalance := bk.GetBalance(ctx, ak.GetModuleAddress(types.BondedPoolName), bondDenom)

	err := bk.SendCoinsFromModuleToModule(ctx, types.BondedPoolName, minttypes.ModuleName, sdk.NewCoins(initialBondedBalance))
	require.NoError(t, err, "no error expected when clearing bonded pool balance")
}

// Helper function to fund the Bonded pool balances before a unit test
func fundPoolBalance(t *testing.T, sk keeper.Keeper, bk bankkeeper.Keeper, ctx sdk.Context, amount sdk.Int) {
	bondDenom := sk.BondDenom(ctx)
	bondedPoolCoin := sdk.NewCoin(bondDenom, amount)

	err := bk.MintCoins(ctx, minttypes.ModuleName, sdk.NewCoins(bondedPoolCoin))
	require.NoError(t, err, "no error expected when minting")

	err = bk.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, types.BondedPoolName, sdk.NewCoins(bondedPoolCoin))
	require.NoError(t, err, "no error expected when sending tokens to bonded pool")
}

// Tests CheckExceedsGlobalLiquidStakingCap
func TestCheckExceedsGlobalLiquidStakingCap(t *testing.T) {
	var (
		accountKeeper accountkeeper.AccountKeeper
		bankKeeper    bankkeeper.Keeper
		stakingKeeper *keeper.Keeper
	)

	app, err := simtestutil.Setup(testutil.AppConfig,
		&accountKeeper,
		&bankKeeper,
		&stakingKeeper,
	)
	require.NoError(t, err)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	testCases := []struct {
		name             string
		globalLiquidCap  sdk.Dec
		totalLiquidStake sdk.Int
		totalStake       sdk.Int
		newLiquidStake   sdk.Int
		tokenizingShares bool
		expectedExceeds  bool
	}{
		{
			// Cap: 10% - Native Delegation - Delegation Below Threshold
			// Total Liquid Stake: 5, Total Stake: 95, New Liquid Stake: 1
			// => Total Liquid Stake: 5+1=6, Total Stake: 95+1=96 => 6/96 = 6% < 10% cap
			name:             "10 percent cap _ native delegation _ delegation below cap",
			globalLiquidCap:  sdk.MustNewDecFromStr("0.1"),
			totalLiquidStake: sdk.NewInt(5),
			totalStake:       sdk.NewInt(95),
			newLiquidStake:   sdk.NewInt(1),
			tokenizingShares: false,
			expectedExceeds:  false,
		},
		{
			// Cap: 10% - Native Delegation - Delegation At Threshold
			// Total Liquid Stake: 5, Total Stake: 95, New Liquid Stake: 5
			// => Total Liquid Stake: 5+5=10, Total Stake: 95+5=100 => 10/100 = 10% == 10% cap
			name:             "10 percent cap _ native delegation _ delegation equals cap",
			globalLiquidCap:  sdk.MustNewDecFromStr("0.1"),
			totalLiquidStake: sdk.NewInt(5),
			totalStake:       sdk.NewInt(95),
			newLiquidStake:   sdk.NewInt(5),
			tokenizingShares: false,
			expectedExceeds:  false,
		},
		{
			// Cap: 10% - Native Delegation - Delegation Exceeds Threshold
			// Total Liquid Stake: 5, Total Stake: 95, New Liquid Stake: 6
			// => Total Liquid Stake: 5+6=11, Total Stake: 95+6=101 => 11/101 = 11% > 10% cap
			name:             "10 percent cap _ native delegation _ delegation exceeds cap",
			globalLiquidCap:  sdk.MustNewDecFromStr("0.1"),
			totalLiquidStake: sdk.NewInt(5),
			totalStake:       sdk.NewInt(95),
			newLiquidStake:   sdk.NewInt(6),
			tokenizingShares: false,
			expectedExceeds:  true,
		},
		{
			// Cap: 20% - Native Delegation - Delegation Below Threshold
			// Total Liquid Stake: 20, Total Stake: 220, New Liquid Stake: 29
			// => Total Liquid Stake: 20+29=49, Total Stake: 220+29=249 => 49/249 = 19% < 20% cap
			name:             "20 percent cap _ native delegation _ delegation below cap",
			globalLiquidCap:  sdk.MustNewDecFromStr("0.20"),
			totalLiquidStake: sdk.NewInt(20),
			totalStake:       sdk.NewInt(220),
			newLiquidStake:   sdk.NewInt(29),
			tokenizingShares: false,
			expectedExceeds:  false,
		},
		{
			// Cap: 20% - Native Delegation - Delegation At Threshold
			// Total Liquid Stake: 20, Total Stake: 220, New Liquid Stake: 30
			// => Total Liquid Stake: 20+30=50, Total Stake: 220+30=250 => 50/250 = 20% == 20% cap
			name:             "20 percent cap _ native delegation _ delegation equals cap",
			globalLiquidCap:  sdk.MustNewDecFromStr("0.20"),
			totalLiquidStake: sdk.NewInt(20),
			totalStake:       sdk.NewInt(220),
			newLiquidStake:   sdk.NewInt(30),
			tokenizingShares: false,
			expectedExceeds:  false,
		},
		{
			// Cap: 20% - Native Delegation - Delegation Exceeds Threshold
			// Total Liquid Stake: 20, Total Stake: 220, New Liquid Stake: 31
			// => Total Liquid Stake: 20+31=51, Total Stake: 220+31=251 => 51/251 = 21% > 20% cap
			name:             "20 percent cap _ native delegation _ delegation exceeds cap",
			globalLiquidCap:  sdk.MustNewDecFromStr("0.20"),
			totalLiquidStake: sdk.NewInt(20),
			totalStake:       sdk.NewInt(220),
			newLiquidStake:   sdk.NewInt(31),
			tokenizingShares: false,
			expectedExceeds:  true,
		},
		{
			// Cap: 50% - Native Delegation - Delegation Below Threshold
			// Total Liquid Stake: 0, Total Stake: 100, New Liquid Stake: 50
			// => Total Liquid Stake: 0+50=50, Total Stake: 100+50=150 => 50/150 = 33% < 50% cap
			name:             "50 percent cap _ native delegation _ delegation below cap",
			globalLiquidCap:  sdk.MustNewDecFromStr("0.5"),
			totalLiquidStake: sdk.NewInt(0),
			totalStake:       sdk.NewInt(100),
			newLiquidStake:   sdk.NewInt(50),
			tokenizingShares: false,
			expectedExceeds:  false,
		},
		{
			// Cap: 50% - Tokenized Delegation - Delegation At Threshold
			// Total Liquid Stake: 0, Total Stake: 100, New Liquid Stake: 50
			// => 50 / 100 = 50% == 50% cap
			name:             "50 percent cap _ tokenized delegation _ delegation equals cap",
			globalLiquidCap:  sdk.MustNewDecFromStr("0.5"),
			totalLiquidStake: sdk.NewInt(0),
			totalStake:       sdk.NewInt(100),
			newLiquidStake:   sdk.NewInt(50),
			tokenizingShares: true,
			expectedExceeds:  false,
		},
		{
			// Cap: 50% - Native Delegation - Delegation Below Threshold
			// Total Liquid Stake: 0, Total Stake: 100, New Liquid Stake: 51
			// => Total Liquid Stake: 0+51=51, Total Stake: 100+51=151 => 51/151 = 33% < 50% cap
			name:             "50 percent cap _ native delegation _ delegation below cap",
			globalLiquidCap:  sdk.MustNewDecFromStr("0.5"),
			totalLiquidStake: sdk.NewInt(0),
			totalStake:       sdk.NewInt(100),
			newLiquidStake:   sdk.NewInt(51),
			tokenizingShares: false,
			expectedExceeds:  false,
		},
		{
			// Cap: 50% - Tokenized Delegation - Delegation Exceeds Threshold
			// Total Liquid Stake: 0, Total Stake: 100, New Liquid Stake: 51
			// => 51 / 100 = 51% > 50% cap
			name:             "50 percent cap _  tokenized delegation _delegation exceeds cap",
			globalLiquidCap:  sdk.MustNewDecFromStr("0.5"),
			totalLiquidStake: sdk.NewInt(0),
			totalStake:       sdk.NewInt(100),
			newLiquidStake:   sdk.NewInt(51),
			tokenizingShares: true,
			expectedExceeds:  true,
		},
		{
			// Cap of 0% - everything should exceed
			name:             "0 percent cap",
			globalLiquidCap:  sdk.ZeroDec(),
			totalLiquidStake: sdk.NewInt(0),
			totalStake:       sdk.NewInt(1_000_000),
			newLiquidStake:   sdk.NewInt(1),
			tokenizingShares: false,
			expectedExceeds:  true,
		},
		{
			// Cap of 100% - nothing should exceed
			name:             "100 percent cap",
			globalLiquidCap:  sdk.OneDec(),
			totalLiquidStake: sdk.NewInt(1),
			totalStake:       sdk.NewInt(1),
			newLiquidStake:   sdk.NewInt(1_000_000),
			tokenizingShares: false,
			expectedExceeds:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Update the global liquid staking cap
			params := stakingKeeper.GetParams(ctx)
			params.GlobalLiquidStakingCap = tc.globalLiquidCap
			stakingKeeper.SetParams(ctx, params)

			// Update the total liquid tokens
			stakingKeeper.SetTotalLiquidStakedTokens(ctx, tc.totalLiquidStake)

			// Fund each pool for the given test case
			clearPoolBalance(t, *stakingKeeper, accountKeeper, bankKeeper, ctx)
			fundPoolBalance(t, *stakingKeeper, bankKeeper, ctx, tc.totalStake)

			// Check if the new tokens would exceed the global cap
			actualExceeds := stakingKeeper.CheckExceedsGlobalLiquidStakingCap(ctx, tc.newLiquidStake, tc.tokenizingShares)
			require.Equal(t, tc.expectedExceeds, actualExceeds, tc.name)
		})
	}
}

// Tests SafelyIncreaseTotalLiquidStakedTokens
func TestSafelyIncreaseTotalLiquidStakedTokens(t *testing.T) {
	var (
		accountKeeper accountkeeper.AccountKeeper
		bankKeeper    bankkeeper.Keeper
		stakingKeeper *keeper.Keeper
	)

	app, err := simtestutil.Setup(testutil.AppConfig,
		&accountKeeper,
		&bankKeeper,
		&stakingKeeper,
	)
	require.NoError(t, err)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	intitialTotalLiquidStaked := sdk.NewInt(100)
	increaseAmount := sdk.NewInt(10)
	poolBalance := sdk.NewInt(200)

	// Set the total staked and total liquid staked amounts
	// which are required components when checking the global cap
	// Total stake is calculated from the pool balance
	clearPoolBalance(t, *stakingKeeper, accountKeeper, bankKeeper, ctx)
	fundPoolBalance(t, *stakingKeeper, bankKeeper, ctx, poolBalance)
	stakingKeeper.SetTotalLiquidStakedTokens(ctx, intitialTotalLiquidStaked)

	// Set the global cap such that a small delegation would exceed the cap
	params := stakingKeeper.GetParams(ctx)
	params.GlobalLiquidStakingCap = sdk.MustNewDecFromStr("0.0001")
	stakingKeeper.SetParams(ctx, params)

	// Attempt to increase the total liquid stake again, it should error since
	// the cap was exceeded
	err = stakingKeeper.SafelyIncreaseTotalLiquidStakedTokens(ctx, increaseAmount, true)
	require.ErrorIs(t, err, types.ErrGlobalLiquidStakingCapExceeded)
	require.Equal(t, intitialTotalLiquidStaked, stakingKeeper.GetTotalLiquidStakedTokens(ctx))

	// Now relax the cap so that the increase succeeds
	params.GlobalLiquidStakingCap = sdk.MustNewDecFromStr("0.99")
	stakingKeeper.SetParams(ctx, params)

	// Confirm the total increased
	err = stakingKeeper.SafelyIncreaseTotalLiquidStakedTokens(ctx, increaseAmount, true)
	require.NoError(t, err)
	require.Equal(t, intitialTotalLiquidStaked.Add(increaseAmount), stakingKeeper.GetTotalLiquidStakedTokens(ctx))
}

// Helper function to create a base account from an account name
// Used to differentiate against liquid staking provider module account
func createBaseAccount(ak accountkeeper.AccountKeeper, ctx sdk.Context, accountName string) sdk.AccAddress {
	baseAccountAddress := sdk.AccAddress(accountName)
	ak.SetAccount(ctx, authtypes.NewBaseAccountWithAddress(baseAccountAddress))
	return baseAccountAddress
}

// Tests DelegatorIsLiquidStaker
func TestDelegatorIsLiquidStaker(t *testing.T) {
	var (
		accountKeeper accountkeeper.AccountKeeper
		stakingKeeper *keeper.Keeper
	)

	app, err := simtestutil.Setup(testutil.AppConfig,
		&accountKeeper,
		&stakingKeeper,
	)
	require.NoError(t, err)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// Create base and ICA accounts
	baseAccountAddress := createBaseAccount(accountKeeper, ctx, "base-account")
	icaAccountAddress := createICAAccount(ctx, accountKeeper)

	// Only the ICA module account should be considered a liquid staking provider
	require.False(t, stakingKeeper.DelegatorIsLiquidStaker(baseAccountAddress), "base account")
	require.True(t, stakingKeeper.DelegatorIsLiquidStaker(icaAccountAddress), "ICA module account")
}

// Tests Add/Remove/Get/SetTokenizeSharesLock
func TestTokenizeSharesLock(t *testing.T) {
	var (
		bankKeeper    bankkeeper.Keeper
		stakingKeeper *keeper.Keeper
	)

	app, err := simtestutil.Setup(testutil.AppConfig,
		&bankKeeper,
		&stakingKeeper,
	)
	require.NoError(t, err)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addresses := simtestutil.AddTestAddrs(bankKeeper, stakingKeeper, ctx, 2, sdk.NewInt(1))
	addressA, addressB := addresses[0], addresses[1]

	unlocked := types.TOKENIZE_SHARE_LOCK_STATUS_UNLOCKED.String()
	locked := types.TOKENIZE_SHARE_LOCK_STATUS_LOCKED.String()
	lockExpiring := types.TOKENIZE_SHARE_LOCK_STATUS_LOCK_EXPIRING.String()

	// Confirm both accounts start unlocked
	status, _ := stakingKeeper.GetTokenizeSharesLock(ctx, addressA)
	require.Equal(t, unlocked, status.String(), "addressA unlocked at start")

	status, _ = stakingKeeper.GetTokenizeSharesLock(ctx, addressB)
	require.Equal(t, unlocked, status.String(), "addressB unlocked at start")

	// Lock the first account
	stakingKeeper.AddTokenizeSharesLock(ctx, addressA)

	// The first account should now have tokenize shares disabled
	// and the unlock time should be the zero time
	status, _ = stakingKeeper.GetTokenizeSharesLock(ctx, addressA)
	require.Equal(t, locked, status.String(), "addressA locked")

	status, _ = stakingKeeper.GetTokenizeSharesLock(ctx, addressB)
	require.Equal(t, unlocked, status.String(), "addressB still unlocked")

	// Update the lock time and confirm it was set
	expectedUnlockTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	stakingKeeper.SetTokenizeSharesUnlockTime(ctx, addressA, expectedUnlockTime)

	status, actualUnlockTime := stakingKeeper.GetTokenizeSharesLock(ctx, addressA)
	require.Equal(t, lockExpiring, status.String(), "addressA lock expiring")
	require.Equal(t, expectedUnlockTime, actualUnlockTime, "addressA unlock time")

	// Confirm B is still unlocked
	status, _ = stakingKeeper.GetTokenizeSharesLock(ctx, addressB)
	require.Equal(t, unlocked, status.String(), "addressB still unlocked")

	// Remove the lock
	stakingKeeper.RemoveTokenizeSharesLock(ctx, addressA)
	status, _ = stakingKeeper.GetTokenizeSharesLock(ctx, addressA)
	require.Equal(t, unlocked, status.String(), "addressA unlocked at end")

	status, _ = stakingKeeper.GetTokenizeSharesLock(ctx, addressB)
	require.Equal(t, unlocked, status.String(), "addressB unlocked at end")
}

// Tests GetAllTokenizeSharesLocks
func TestGetAllTokenizeSharesLocks(t *testing.T) {
	var (
		bankKeeper    bankkeeper.Keeper
		stakingKeeper *keeper.Keeper
	)

	app, err := simtestutil.Setup(testutil.AppConfig,
		&bankKeeper,
		&stakingKeeper,
	)
	require.NoError(t, err)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addresses := simtestutil.AddTestAddrs(bankKeeper, stakingKeeper, ctx, 4, sdk.NewInt(1))

	// Set 2 locked accounts, and two accounts with a lock expiring
	stakingKeeper.AddTokenizeSharesLock(ctx, addresses[0])
	stakingKeeper.AddTokenizeSharesLock(ctx, addresses[1])

	unlockTime1 := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
	unlockTime2 := time.Date(2023, 1, 2, 1, 0, 0, 0, time.UTC)
	stakingKeeper.SetTokenizeSharesUnlockTime(ctx, addresses[2], unlockTime1)
	stakingKeeper.SetTokenizeSharesUnlockTime(ctx, addresses[3], unlockTime2)

	// Defined expected locks after GetAll
	expectedLocks := map[string]types.TokenizeShareLock{
		addresses[0].String(): {
			Status: types.TOKENIZE_SHARE_LOCK_STATUS_LOCKED.String(),
		},
		addresses[1].String(): {
			Status: types.TOKENIZE_SHARE_LOCK_STATUS_LOCKED.String(),
		},
		addresses[2].String(): {
			Status:         types.TOKENIZE_SHARE_LOCK_STATUS_LOCK_EXPIRING.String(),
			CompletionTime: unlockTime1,
		},
		addresses[3].String(): {
			Status:         types.TOKENIZE_SHARE_LOCK_STATUS_LOCK_EXPIRING.String(),
			CompletionTime: unlockTime2,
		},
	}

	// Check output from GetAll
	actualLocks := stakingKeeper.GetAllTokenizeSharesLocks(ctx)
	require.Len(t, actualLocks, len(expectedLocks), "number of locks")

	for i, actual := range actualLocks {
		expected, ok := expectedLocks[actual.Address]
		require.True(t, ok, "address %s not expected", actual.Address)
		require.Equal(t, expected.Status, actual.Status, "tokenize share lock #%d status", i)
		require.Equal(t, expected.CompletionTime, actual.CompletionTime, "tokenize share lock #%d completion time", i)
	}
}

// Test Get/SetPendingTokenizeShareAuthorizations
func TestPendingTokenizeShareAuthorizations(t *testing.T) {
	var (
		bankKeeper    bankkeeper.Keeper
		stakingKeeper *keeper.Keeper
	)

	app, err := simtestutil.Setup(testutil.AppConfig,
		&bankKeeper,
		&stakingKeeper,
	)
	require.NoError(t, err)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// Create dummy accounts and completion times
	addresses := simtestutil.AddTestAddrs(bankKeeper, stakingKeeper, ctx, 3, sdk.NewInt(1))
	addressStrings := []string{}
	for _, address := range addresses {
		addressStrings = append(addressStrings, address.String())
	}

	timeA := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	timeB := timeA.Add(time.Hour)

	// There should be no addresses returned originally
	authorizationsA := stakingKeeper.GetPendingTokenizeShareAuthorizations(ctx, timeA)
	require.Empty(t, authorizationsA.Addresses, "no addresses at timeA expected")

	authorizationsB := stakingKeeper.GetPendingTokenizeShareAuthorizations(ctx, timeB)
	require.Empty(t, authorizationsB.Addresses, "no addresses at timeB expected")

	// Store addresses for timeB
	stakingKeeper.SetPendingTokenizeShareAuthorizations(ctx, timeB, types.PendingTokenizeShareAuthorizations{
		Addresses: addressStrings,
	})

	// Check addresses
	authorizationsA = stakingKeeper.GetPendingTokenizeShareAuthorizations(ctx, timeA)
	require.Empty(t, authorizationsA.Addresses, "no addresses at timeA expected at end")

	authorizationsB = stakingKeeper.GetPendingTokenizeShareAuthorizations(ctx, timeB)
	require.Equal(t, addressStrings, authorizationsB.Addresses, "address length")
}

// Test QueueTokenizeSharesAuthorization and RemoveExpiredTokenizeShareLocks
func TestTokenizeShareAuthorizationQueue(t *testing.T) {
	var (
		bankKeeper    bankkeeper.Keeper
		stakingKeeper *keeper.Keeper
	)

	app, err := simtestutil.Setup(testutil.AppConfig,
		&bankKeeper,
		&stakingKeeper,
	)
	require.NoError(t, err)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// We'll start by adding the following addresses to the queue
	//   Time 0: [address0]
	//   Time 1: []
	//   Time 2: [address1, address2, address3]
	//   Time 3: [address4, address5]
	//   Time 4: [address6]
	addresses := simtestutil.AddTestAddrs(bankKeeper, stakingKeeper, ctx, 7, sdk.NewInt(1))
	addressesByTime := map[int][]sdk.AccAddress{
		0: {addresses[0]},
		1: {},
		2: {addresses[1], addresses[2], addresses[3]},
		3: {addresses[4], addresses[5]},
		4: {addresses[6]},
	}

	// Set the unbonding time to 1 day
	unbondingPeriod := time.Hour * 24
	params := stakingKeeper.GetParams(ctx)
	params.UnbondingTime = unbondingPeriod
	stakingKeeper.SetParams(ctx, params)

	// Add each address to the queue and then increment the block time
	// such that the times line up as follows
	//   Time 0: 2023-01-01 00:00:00
	//   Time 1: 2023-01-01 00:01:00
	//   Time 2: 2023-01-01 00:02:00
	//   Time 3: 2023-01-01 00:03:00
	startTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	ctx = ctx.WithBlockTime(startTime)
	blockTimeIncrement := time.Hour

	for timeIndex := 0; timeIndex <= 4; timeIndex++ {
		for _, address := range addressesByTime[timeIndex] {
			stakingKeeper.QueueTokenizeSharesAuthorization(ctx, address)
		}
		ctx = ctx.WithBlockTime(ctx.BlockTime().Add(blockTimeIncrement))
	}

	// We'll unlock the tokens using the following progression
	// The "alias'"/keys for these times assume a starting point of the Time 0
	// from above, plus the Unbonding Time
	//   Time -1  (2023-01-01 23:59:99): []
	//   Time  0  (2023-01-02 00:00:00): [address0]
	//   Time  1  (2023-01-02 00:01:00): []
	//   Time 2.5 (2023-01-02 00:02:30): [address1, address2, address3]
	//   Time 10  (2023-01-02 00:10:00): [address4, address5, address6]
	unlockBlockTimes := map[string]time.Time{
		"-1":  startTime.Add(unbondingPeriod).Add(-time.Second),
		"0":   startTime.Add(unbondingPeriod),
		"1":   startTime.Add(unbondingPeriod).Add(blockTimeIncrement),
		"2.5": startTime.Add(unbondingPeriod).Add(2 * blockTimeIncrement).Add(blockTimeIncrement / 2),
		"10":  startTime.Add(unbondingPeriod).Add(10 * blockTimeIncrement),
	}
	expectedUnlockedAddresses := map[string][]string{
		"-1":  {},
		"0":   {addresses[0].String()},
		"1":   {},
		"2.5": {addresses[1].String(), addresses[2].String(), addresses[3].String()},
		"10":  {addresses[4].String(), addresses[5].String(), addresses[6].String()},
	}

	// Now we'll remove items from the queue sequentially
	// First check with a block time before the first expiration - it should remove no addresses
	actualAddresses := stakingKeeper.RemoveExpiredTokenizeShareLocks(ctx, unlockBlockTimes["-1"])
	require.Equal(t, expectedUnlockedAddresses["-1"], actualAddresses, "no addresses unlocked from time -1")

	// Then pass in (time 0 + unbonding time) - it should remove the first address
	actualAddresses = stakingKeeper.RemoveExpiredTokenizeShareLocks(ctx, unlockBlockTimes["0"])
	require.Equal(t, expectedUnlockedAddresses["0"], actualAddresses, "one address unlocked from time 0")

	// Now pass in (time 1 + unbonding time) - it should remove no addresses since
	// the address at time 0 was already removed
	actualAddresses = stakingKeeper.RemoveExpiredTokenizeShareLocks(ctx, unlockBlockTimes["1"])
	require.Equal(t, expectedUnlockedAddresses["1"], actualAddresses, "no addresses unlocked from time 1")

	// Now pass in (time 2.5 + unbonding time) - it should remove the three addresses from time 2
	actualAddresses = stakingKeeper.RemoveExpiredTokenizeShareLocks(ctx, unlockBlockTimes["2.5"])
	require.Equal(t, expectedUnlockedAddresses["2.5"], actualAddresses, "addresses unlocked from time 2.5")

	// Finally pass in a block time far in the future, which should remove all the remaining locks
	actualAddresses = stakingKeeper.RemoveExpiredTokenizeShareLocks(ctx, unlockBlockTimes["10"])
	require.Equal(t, expectedUnlockedAddresses["10"], actualAddresses, "addresses unlocked from time 10")
}
