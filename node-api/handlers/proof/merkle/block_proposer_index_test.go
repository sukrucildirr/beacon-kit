// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package merkle_test

import (
	"testing"

	"github.com/berachain/beacon-kit/consensus-types/types"
	"github.com/berachain/beacon-kit/node-api/handlers/proof/merkle"
	"github.com/berachain/beacon-kit/primitives/common"
	"github.com/berachain/beacon-kit/primitives/math"
	"github.com/stretchr/testify/require"
)

// TestBlockProposerIndexProof tests the ProveProposerIndexInBlock function
// and that the generated proof correctly verifies.
func TestBlockProposerIndexProof(t *testing.T) {
	testCases := []struct {
		name              string
		slot              math.Slot
		proposerIndex     math.ValidatorIndex
		parentBlockRoot   common.Root
		stateRoot         common.Root
		bodyRoot          common.Root
		expectedProofFile string
	}{
		{
			name:              "1 Validator Set",
			slot:              69,
			proposerIndex:     0,
			parentBlockRoot:   common.Root{1, 2, 3},
			stateRoot:         common.Root{4, 5, 6},
			bodyRoot:          common.Root{7, 8, 9},
			expectedProofFile: "one_validator_proposer_index_proof.json",
		},
		{
			name:              "Many Validator Set",
			slot:              420,
			proposerIndex:     69,
			parentBlockRoot:   common.Root{1, 2, 3},
			stateRoot:         common.Root{4, 5, 6},
			bodyRoot:          common.Root{7, 8, 9},
			expectedProofFile: "many_validators_proposer_index_proof.json",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bbh := types.NewBeaconBlockHeader(
				tc.slot,
				tc.proposerIndex,
				tc.parentBlockRoot,
				tc.stateRoot,
				tc.bodyRoot,
			)

			proof, beaconRoot, err := merkle.ProveProposerIndexInBlock(bbh)
			require.NoError(t, err)

			require.Equal(t, bbh.HashTreeRoot(), beaconRoot)

			expectedProof := ReadProofFromFile(t, tc.expectedProofFile)
			require.Equal(t, expectedProof, proof)
		})
	}
}
