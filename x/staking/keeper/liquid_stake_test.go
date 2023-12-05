package keeper_test

import (
	"fmt"

	"cosmossdk.io/simapp"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Helper function to create a base account from an account name
// Used to differentiate against liquid staking provider module account
func createBaseAccount(app *simapp.SimApp, ctx sdk.Context, accountName string) sdk.AccAddress {
	baseAccountAddress := sdk.AccAddress(accountName)
	app.AccountKeeper.SetAccount(ctx, authtypes.NewBaseAccountWithAddress(baseAccountAddress))
	return baseAccountAddress
}

// Helper function to create a module account address from a tokenized share
// Used to mock the delegation owner of a tokenized share
func createTokenizeShareModuleAccount(recordID uint64) sdk.AccAddress {
	record := types.TokenizeShareRecord{
		Id:            recordID,
		ModuleAccount: fmt.Sprintf("%s%d", types.TokenizeShareModuleAccountPrefix, recordID),
	}
	return record.GetModuleAddress()
}

// Tests Set/Get TotalLiquidStakedTokens
func (s *KeeperTestSuite) TestTotalLiquidStakedTokens() {
	ctx, keeper := s.ctx, s.stakingKeeper
	require := s.Require()

	// Update the total liquid staked
	total := sdk.NewInt(100)
	keeper.SetTotalLiquidStakedTokens(ctx, total)

	// Confirm it was updated
	require.Equal(total, keeper.GetTotalLiquidStakedTokens(ctx), "initial")
}

// Tests Increase/Decrease TotalValidatorLiquidShares
func (s *KeeperTestSuite) TestValidatorLiquidShares() {
	ctx, keeper := s.ctx, s.stakingKeeper

	// Create a validator address
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	valAddress := sdk.ValAddress(pubKey.Address())

	// Set an initial total
	initial := sdk.NewDec(100)
	validator := types.Validator{
		OperatorAddress: valAddress.String(),
		LiquidShares:    initial,
	}
	keeper.SetValidator(ctx, validator)
}

// Tests DecreaseTotalLiquidStakedTokens
func (s *KeeperTestSuite) TestDecreaseTotalLiquidStakedTokens() {
	ctx, keeper := s.ctx, s.stakingKeeper
	require := s.Require()

	intitialTotalLiquidStaked := sdk.NewInt(100)
	decreaseAmount := sdk.NewInt(10)

	// Set the total liquid staked to an arbitrary value
	keeper.SetTotalLiquidStakedTokens(ctx, intitialTotalLiquidStaked)

	// Decrease the total liquid stake and confirm the total was updated
	err := keeper.DecreaseTotalLiquidStakedTokens(ctx, decreaseAmount)
	require.NoError(err, "no error expected when decreasing total liquid staked tokens")
	require.Equal(intitialTotalLiquidStaked.Sub(decreaseAmount), keeper.GetTotalLiquidStakedTokens(ctx))

	// Attempt to decrease by an excessive amount, it should error
	err = keeper.DecreaseTotalLiquidStakedTokens(ctx, intitialTotalLiquidStaked)
	require.ErrorIs(err, types.ErrTotalLiquidStakedUnderflow)
}

// Tests CheckExceedsValidatorBondCap
func (s *KeeperTestSuite) TestCheckExceedsValidatorBondCap() {
	ctx, keeper := s.ctx, s.stakingKeeper
	require := s.Require()

	testCases := []struct {
		name                string
		validatorShares     sdk.Dec
		validatorBondFactor sdk.Dec
		currentLiquidShares sdk.Dec
		newShares           sdk.Dec
		expectedExceeds     bool
	}{
		{
			// Validator Shares: 100, Factor: 1, Current Shares: 90 => 100 Max Shares, Capacity: 10
			// New Shares: 5 - below cap
			name:                "factor 1 - below cap",
			validatorShares:     sdk.NewDec(100),
			validatorBondFactor: sdk.NewDec(1),
			currentLiquidShares: sdk.NewDec(90),
			newShares:           sdk.NewDec(5),
			expectedExceeds:     false,
		},
		{
			// Validator Shares: 100, Factor: 1, Current Shares: 90 => 100 Max Shares, Capacity: 10
			// New Shares: 10 - at cap
			name:                "factor 1 - at cap",
			validatorShares:     sdk.NewDec(100),
			validatorBondFactor: sdk.NewDec(1),
			currentLiquidShares: sdk.NewDec(90),
			newShares:           sdk.NewDec(10),
			expectedExceeds:     false,
		},
		{
			// Validator Shares: 100, Factor: 1, Current Shares: 90 => 100 Max Shares, Capacity: 10
			// New Shares: 15 - above cap
			name:                "factor 1 - above cap",
			validatorShares:     sdk.NewDec(100),
			validatorBondFactor: sdk.NewDec(1),
			currentLiquidShares: sdk.NewDec(90),
			newShares:           sdk.NewDec(15),
			expectedExceeds:     true,
		},
		{
			// Validator Shares: 100, Factor: 2, Current Shares: 90 => 200 Max Shares, Capacity: 110
			// New Shares: 5 - below cap
			name:                "factor 2 - well below cap",
			validatorShares:     sdk.NewDec(100),
			validatorBondFactor: sdk.NewDec(2),
			currentLiquidShares: sdk.NewDec(90),
			newShares:           sdk.NewDec(5),
			expectedExceeds:     false,
		},
		{
			// Validator Shares: 100, Factor: 2, Current Shares: 90 => 200 Max Shares, Capacity: 110
			// New Shares: 100 - below cap
			name:                "factor 2 - below cap",
			validatorShares:     sdk.NewDec(100),
			validatorBondFactor: sdk.NewDec(2),
			currentLiquidShares: sdk.NewDec(90),
			newShares:           sdk.NewDec(100),
			expectedExceeds:     false,
		},
		{
			// Validator Shares: 100, Factor: 2, Current Shares: 90 => 200 Max Shares, Capacity: 110
			// New Shares: 110 - below cap
			name:                "factor 2 - at cap",
			validatorShares:     sdk.NewDec(100),
			validatorBondFactor: sdk.NewDec(2),
			currentLiquidShares: sdk.NewDec(90),
			newShares:           sdk.NewDec(110),
			expectedExceeds:     false,
		},
		{
			// Validator Shares: 100, Factor: 2, Current Shares: 90 => 200 Max Shares, Capacity: 110
			// New Shares: 111 - above cap
			name:                "factor 2 - above cap",
			validatorShares:     sdk.NewDec(100),
			validatorBondFactor: sdk.NewDec(2),
			currentLiquidShares: sdk.NewDec(90),
			newShares:           sdk.NewDec(111),
			expectedExceeds:     true,
		},
		{
			// Validator Shares: 100, Factor: 100, Current Shares: 90 => 10000 Max Shares, Capacity: 9910
			// New Shares: 100 - below cap
			name:                "factor 100 - below cap",
			validatorShares:     sdk.NewDec(100),
			validatorBondFactor: sdk.NewDec(100),
			currentLiquidShares: sdk.NewDec(90),
			newShares:           sdk.NewDec(100),
			expectedExceeds:     false,
		},
		{
			// Validator Shares: 100, Factor: 100, Current Shares: 90 => 10000 Max Shares, Capacity: 9910
			// New Shares: 9910 - at cap
			name:                "factor 100 - at cap",
			validatorShares:     sdk.NewDec(100),
			validatorBondFactor: sdk.NewDec(100),
			currentLiquidShares: sdk.NewDec(90),
			newShares:           sdk.NewDec(9910),
			expectedExceeds:     false,
		},
		{
			// Validator Shares: 100, Factor: 100, Current Shares: 90 => 10000 Max Shares, Capacity: 9910
			// New Shares: 9911 - above cap
			name:                "factor 100 - above cap",
			validatorShares:     sdk.NewDec(100),
			validatorBondFactor: sdk.NewDec(100),
			currentLiquidShares: sdk.NewDec(90),
			newShares:           sdk.NewDec(9911),
			expectedExceeds:     true,
		},
		{
			// Factor of -1 (disabled): Should always return false
			name:                "factor disabled",
			validatorShares:     sdk.NewDec(1),
			validatorBondFactor: sdk.NewDec(-1),
			currentLiquidShares: sdk.NewDec(1),
			newShares:           sdk.NewDec(1_000_000),
			expectedExceeds:     false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Update the validator bond factor
			params := keeper.GetParams(ctx)
			params.ValidatorBondFactor = tc.validatorBondFactor
			keeper.SetParams(ctx, params)

			// Create a validator with designated self-bond shares
			validator := types.Validator{
				LiquidShares:        tc.currentLiquidShares,
				ValidatorBondShares: tc.validatorShares,
			}

			// Check whether the cap is exceeded
			actualExceeds := keeper.CheckExceedsValidatorBondCap(ctx, validator, tc.newShares)
			require.Equal(tc.expectedExceeds, actualExceeds, tc.name)
		})
	}
}

// Tests TestCheckExceedsValidatorLiquidStakingCap
func (s *KeeperTestSuite) TestCheckExceedsValidatorLiquidStakingCap() {
	ctx, keeper := s.ctx, s.stakingKeeper
	require := s.Require()

	testCases := []struct {
		name                  string
		validatorLiquidCap    sdk.Dec
		validatorLiquidShares sdk.Dec
		validatorTotalShares  sdk.Dec
		newLiquidShares       sdk.Dec
		expectedExceeds       bool
	}{
		{
			// Cap: 10% - Delegation Below Threshold
			// Liquid Shares: 5, Total Shares: 95, New Liquid Shares: 1
			// => Liquid Shares: 5+1=6, Total Shares: 95+1=96 => 6/96 = 6% < 10% cap
			name:                  "10 percent cap _ delegation below cap",
			validatorLiquidCap:    sdk.MustNewDecFromStr("0.1"),
			validatorLiquidShares: sdk.NewDec(5),
			validatorTotalShares:  sdk.NewDec(95),
			newLiquidShares:       sdk.NewDec(1),
			expectedExceeds:       false,
		},
		{
			// Cap: 10% - Delegation At Threshold
			// Liquid Shares: 5, Total Shares: 95, New Liquid Shares: 5
			// => Liquid Shares: 5+5=10, Total Shares: 95+5=100 => 10/100 = 10% == 10% cap
			name:                  "10 percent cap _ delegation equals cap",
			validatorLiquidCap:    sdk.MustNewDecFromStr("0.1"),
			validatorLiquidShares: sdk.NewDec(5),
			validatorTotalShares:  sdk.NewDec(95),
			newLiquidShares:       sdk.NewDec(4),
			expectedExceeds:       false,
		},
		{
			// Cap: 10% - Delegation Exceeds Threshold
			// Liquid Shares: 5, Total Shares: 95, New Liquid Shares: 6
			// => Liquid Shares: 5+6=11, Total Shares: 95+6=101 => 11/101 = 11% > 10% cap
			name:                  "10 percent cap _ delegation exceeds cap",
			validatorLiquidCap:    sdk.MustNewDecFromStr("0.1"),
			validatorLiquidShares: sdk.NewDec(5),
			validatorTotalShares:  sdk.NewDec(95),
			newLiquidShares:       sdk.NewDec(6),
			expectedExceeds:       true,
		},
		{
			// Cap: 20% - Delegation Below Threshold
			// Liquid Shares: 20, Total Shares: 220, New Liquid Shares: 29
			// => Liquid Shares: 20+29=49, Total Shares: 220+29=249 => 49/249 = 19% < 20% cap
			name:                  "20 percent cap _ delegation below cap",
			validatorLiquidCap:    sdk.MustNewDecFromStr("0.2"),
			validatorLiquidShares: sdk.NewDec(20),
			validatorTotalShares:  sdk.NewDec(220),
			newLiquidShares:       sdk.NewDec(29),
			expectedExceeds:       false,
		},
		{
			// Cap: 20% - Delegation At Threshold
			// Liquid Shares: 20, Total Shares: 220, New Liquid Shares: 30
			// => Liquid Shares: 20+30=50, Total Shares: 220+30=250 => 50/250 = 20% == 20% cap
			name:                  "20 percent cap _ delegation equals cap",
			validatorLiquidCap:    sdk.MustNewDecFromStr("0.2"),
			validatorLiquidShares: sdk.NewDec(20),
			validatorTotalShares:  sdk.NewDec(220),
			newLiquidShares:       sdk.NewDec(30),
			expectedExceeds:       false,
		},
		{
			// Cap: 20% - Delegation Exceeds Threshold
			// Liquid Shares: 20, Total Shares: 220, New Liquid Shares: 31
			// => Liquid Shares: 20+31=51, Total Shares: 220+31=251 => 51/251 = 21% > 20% cap
			name:                  "20 percent cap _ delegation exceeds cap",
			validatorLiquidCap:    sdk.MustNewDecFromStr("0.2"),
			validatorLiquidShares: sdk.NewDec(20),
			validatorTotalShares:  sdk.NewDec(220),
			newLiquidShares:       sdk.NewDec(31),
			expectedExceeds:       true,
		},
		{
			// Cap of 0% - everything should exceed
			name:                  "0 percent cap",
			validatorLiquidCap:    sdk.ZeroDec(),
			validatorLiquidShares: sdk.NewDec(0),
			validatorTotalShares:  sdk.NewDec(1_000_000),
			newLiquidShares:       sdk.NewDec(1),
			expectedExceeds:       true,
		},
		{
			// Cap of 100% - nothing should exceed
			name:                  "100 percent cap",
			validatorLiquidCap:    sdk.OneDec(),
			validatorLiquidShares: sdk.NewDec(1),
			validatorTotalShares:  sdk.NewDec(1_000_000),
			newLiquidShares:       sdk.NewDec(1),
			expectedExceeds:       false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Update the validator liquid staking cap
			params := keeper.GetParams(ctx)
			params.ValidatorLiquidStakingCap = tc.validatorLiquidCap
			keeper.SetParams(ctx, params)

			// Create a validator with designated self-bond shares
			validator := types.Validator{
				LiquidShares:    tc.validatorLiquidShares,
				DelegatorShares: tc.validatorTotalShares,
			}

			// Check whether the cap is exceeded
			actualExceeds := keeper.CheckExceedsValidatorLiquidStakingCap(ctx, validator, tc.newLiquidShares)
			require.Equal(tc.expectedExceeds, actualExceeds, tc.name)
		})
	}
}

// Tests SafelyIncreaseValidatorLiquidShares
func (s *KeeperTestSuite) TestSafelyIncreaseValidatorLiquidShares() {
	ctx, keeper := s.ctx, s.stakingKeeper
	require := s.Require()

	// Generate a test validator address
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	valAddress := sdk.ValAddress(pubKey.Address())

	// Helper function to check the validator's liquid shares
	checkValidatorLiquidShares := func(expected sdk.Dec, description string) {
		actualValidator, found := keeper.GetValidator(ctx, valAddress)
		require.True(found)
		require.Equal(expected.TruncateInt64(), actualValidator.LiquidShares.TruncateInt64(), description)
	}

	// Start with the following:
	//   Initial Liquid Shares: 0
	//   Validator Bond Shares: 10
	//   Validator TotalShares: 75
	//
	// Initial Caps:
	//   ValidatorBondFactor: 1 (Cap applied at 10 shares)
	//   ValidatorLiquidStakingCap: 25% (Cap applied at 25 shares)
	//
	// Cap Increases:
	//   ValidatorBondFactor: 10 (Cap applied at 100 shares)
	//   ValidatorLiquidStakingCap: 40% (Cap applied at 50 shares)
	initialLiquidShares := sdk.NewDec(0)
	validatorBondShares := sdk.NewDec(10)
	validatorTotalShares := sdk.NewDec(75)

	firstIncreaseAmount := sdk.NewDec(20)
	secondIncreaseAmount := sdk.NewDec(10) // total increase of 30

	initialBondFactor := sdk.NewDec(1)
	finalBondFactor := sdk.NewDec(10)
	initialLiquidStakingCap := sdk.MustNewDecFromStr("0.25")
	finalLiquidStakingCap := sdk.MustNewDecFromStr("0.4")

	// Create a validator with designated self-bond shares
	initialValidator := types.Validator{
		OperatorAddress:     valAddress.String(),
		LiquidShares:        initialLiquidShares,
		ValidatorBondShares: validatorBondShares,
		DelegatorShares:     validatorTotalShares,
	}
	keeper.SetValidator(ctx, initialValidator)

	// Set validator bond factor to a small number such that any delegation would fail,
	// and set the liquid staking cap such that the first stake would succeed, but the second
	// would fail
	params := keeper.GetParams(ctx)
	params.ValidatorBondFactor = initialBondFactor
	params.ValidatorLiquidStakingCap = initialLiquidStakingCap
	keeper.SetParams(ctx, params)

	// Attempt to increase the validator liquid shares, it should throw an
	// error that the validator bond cap was exceeded
	_, err := keeper.SafelyIncreaseValidatorLiquidShares(ctx, valAddress, firstIncreaseAmount)
	require.ErrorIs(err, types.ErrInsufficientValidatorBondShares)
	checkValidatorLiquidShares(initialLiquidShares, "shares after low bond factor")

	// Change validator bond factor to a more conservative number, so that the increase succeeds
	params.ValidatorBondFactor = finalBondFactor
	keeper.SetParams(ctx, params)

	// Try the increase again and check that it succeeded
	expectedLiquidSharesAfterFirstStake := initialLiquidShares.Add(firstIncreaseAmount)
	_, err = keeper.SafelyIncreaseValidatorLiquidShares(ctx, valAddress, firstIncreaseAmount)
	require.NoError(err)
	checkValidatorLiquidShares(expectedLiquidSharesAfterFirstStake, "shares with cap loose bond cap")

	// Attempt another increase, it should fail from the liquid staking cap
	_, err = keeper.SafelyIncreaseValidatorLiquidShares(ctx, valAddress, secondIncreaseAmount)
	require.ErrorIs(err, types.ErrValidatorLiquidStakingCapExceeded)
	checkValidatorLiquidShares(expectedLiquidSharesAfterFirstStake, "shares after liquid staking cap hit")

	// Raise the liquid staking cap so the new increment succeeds
	params.ValidatorLiquidStakingCap = finalLiquidStakingCap
	keeper.SetParams(ctx, params)

	// Finally confirm that the increase succeeded this time
	expectedLiquidSharesAfterSecondStake := expectedLiquidSharesAfterFirstStake.Add(secondIncreaseAmount)
	_, err = keeper.SafelyIncreaseValidatorLiquidShares(ctx, valAddress, secondIncreaseAmount)
	require.NoError(err, "no error expected after increasing liquid staking cap")
	checkValidatorLiquidShares(expectedLiquidSharesAfterSecondStake, "shares after loose liquid stake cap")
}

// Tests DecreaseValidatorLiquidShares
func (s *KeeperTestSuite) TestDecreaseValidatorLiquidShares() {
	ctx, keeper := s.ctx, s.stakingKeeper
	require := s.Require()

	initialLiquidShares := sdk.NewDec(100)
	decreaseAmount := sdk.NewDec(10)

	// Create a validator with designated self-bond shares
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	valAddress := sdk.ValAddress(pubKey.Address())

	initialValidator := types.Validator{
		OperatorAddress: valAddress.String(),
		LiquidShares:    initialLiquidShares,
	}
	keeper.SetValidator(ctx, initialValidator)

	// Decrease the validator liquid shares, and confirm the new share amount has been updated
	_, err := keeper.DecreaseValidatorLiquidShares(ctx, valAddress, decreaseAmount)
	require.NoError(err, "no error expected when decreasing validator liquid shares")

	actualValidator, found := keeper.GetValidator(ctx, valAddress)
	require.True(found)
	require.Equal(initialLiquidShares.Sub(decreaseAmount), actualValidator.LiquidShares, "liquid shares")

	// Attempt to decrease by a larger amount than it has, it should fail
	_, err = keeper.DecreaseValidatorLiquidShares(ctx, valAddress, initialLiquidShares)
	require.ErrorIs(err, types.ErrValidatorLiquidSharesUnderflow)
}

// Tests SafelyDecreaseValidatorBond
func (s *KeeperTestSuite) TestSafelyDecreaseValidatorBond() {
	ctx, keeper := s.ctx, s.stakingKeeper
	require := s.Require()

	// Initial Bond Factor: 100, Initial Validator Bond: 10
	// => Max Liquid Shares 1000 (Initial Liquid Shares: 200)
	initialBondFactor := sdk.NewDec(100)
	initialValidatorBondShares := sdk.NewDec(10)
	initialLiquidShares := sdk.NewDec(200)

	// Create a validator with designated self-bond shares
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	valAddress := sdk.ValAddress(pubKey.Address())

	initialValidator := types.Validator{
		OperatorAddress:     valAddress.String(),
		ValidatorBondShares: initialValidatorBondShares,
		LiquidShares:        initialLiquidShares,
	}
	keeper.SetValidator(ctx, initialValidator)

	// Set the bond factor
	params := keeper.GetParams(ctx)
	params.ValidatorBondFactor = initialBondFactor
	keeper.SetParams(ctx, params)

	// Decrease the validator bond from 10 to 5 (minus 5)
	// This will adjust the cap (factor * shares)
	// from (100 * 10 = 1000) to (100 * 5 = 500)
	// Since this is still above the initial liquid shares of 200, this will succeed
	decreaseAmount, expectedBondShares := sdk.NewDec(5), sdk.NewDec(5)
	err := keeper.SafelyDecreaseValidatorBond(ctx, valAddress, decreaseAmount)
	require.NoError(err)

	actualValidator, found := keeper.GetValidator(ctx, valAddress)
	require.True(found)
	require.Equal(expectedBondShares, actualValidator.ValidatorBondShares, "validator bond shares shares")

	// Now attempt to decrease the validator bond again from 5 to 1 (minus 4)
	// This time, the cap will be reduced to (factor * shares) = (100 * 1) = 100
	// However, the liquid shares are currently 200, so this should fail
	decreaseAmount, expectedBondShares = sdk.NewDec(4), sdk.NewDec(1)
	err = keeper.SafelyDecreaseValidatorBond(ctx, valAddress, decreaseAmount)
	require.ErrorIs(err, types.ErrInsufficientValidatorBondShares)

	// Finally, disable the cap and attempt to decrease again
	// This time it should succeed
	params.ValidatorBondFactor = types.ValidatorBondCapDisabled
	keeper.SetParams(ctx, params)

	err = keeper.SafelyDecreaseValidatorBond(ctx, valAddress, decreaseAmount)
	require.NoError(err)

	actualValidator, found = keeper.GetValidator(ctx, valAddress)
	require.True(found)
	require.Equal(expectedBondShares, actualValidator.ValidatorBondShares, "validator bond shares shares")
}
