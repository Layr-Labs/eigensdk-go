// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import {ISlashingRegistryCoordinatorTypes} from "eigenlayer-middleware/src/interfaces/ISlashingRegistryCoordinator.sol";
import {IBLSApkRegistryTypes} from "eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import {ISignatureUtilsMixinTypes} from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtilsMixin.sol";

// Tests are used to verify the encoding of operator registration parameters are correct.
// The results are used in chainio/clients/elcontracts/writer_test.go
contract RegistrationEncodingTest is Test {
    function _getPubkeyRegistrationParams()
        internal
        pure
        returns (IBLSApkRegistryTypes.PubkeyRegistrationParams memory params)
    {
        // Values are random
        params.pubkeyRegistrationSignature.X =
            756874975973566196338995715738218418291193261429375530560923897690728869289;
        params.pubkeyRegistrationSignature.Y =
            444340189040315797681399101731743234568891767085799644128199800550863908703;
        params.pubkeyG1.X = 10371454967541283327403832945957227913391851874635485454053224012738342927470;
        params.pubkeyG1.Y = 5591557118325006940652332791312874698324071372762903093203759620236776485604;
        params.pubkeyG2.X = [
            1357671944470767405259541876666418155809079857448568479000932998534484593852,
            5283708918582394678225755661661470830476341241033602812294790796852421312310
        ];
        params.pubkeyG2.Y = [
            17411007011468414688052176335308121672943440336543041782092111920184779631952,
            6486088401181402728530570019430319265466049090881984123818353102069786218525
        ];
    }

    function testNormalRegistrationEncoding() public pure {
        ISlashingRegistryCoordinatorTypes.RegistrationType registrationType =
            ISlashingRegistryCoordinatorTypes.RegistrationType.NORMAL;
        string memory socket = "unused";
        IBLSApkRegistryTypes.PubkeyRegistrationParams memory params = _getPubkeyRegistrationParams();

        assertEq(
            abi.encode(registrationType, socket, params),
            hex"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014001ac6045296d64b31ed644e53ce1a1c4f72f67a2d47b06b652ca8167f1b2ada900fb7cd59f322f4dffa18360bfcdc15f1f6cd09aea573efa919dc828dfabaf5f16ee091592629fc566636de7b3d53322f4833014b4655dd57279cfb2828bbc6e0c5cb58c8d572dc9dcf5f5999501533e22243fac4f8ee4452a6dd0d4bcaeb2e403006a43453d56eafa7dc4ddcbd41b2330031f58e437ee3806c50a9e554a0cbc0bae792831463d56a1a9983647b77fcdbae38a6621ceef9e147614bb4869bf36267e47def74c144f8e7238dd088097943d1007f8ce7cab028151e9a1beac6d500e56fef5ea67a586d1fbb03dcc0a268f6c4835ad1c215eadc77cc676c378101d0000000000000000000000000000000000000000000000000000000000000006756e757365640000000000000000000000000000000000000000000000000000"
        );
    }

    function testChurnRegistrationEncoding() public pure {
        ISlashingRegistryCoordinatorTypes.RegistrationType registrationType =
            ISlashingRegistryCoordinatorTypes.RegistrationType.CHURN;
        string memory socket = "unused";
        IBLSApkRegistryTypes.PubkeyRegistrationParams memory params = _getPubkeyRegistrationParams();

        // Values are random
        ISlashingRegistryCoordinatorTypes.OperatorKickParam[] memory operatorKickParams =
            new ISlashingRegistryCoordinatorTypes.OperatorKickParam[](2);

        operatorKickParams[0].quorumNumber = 0x0;
        operatorKickParams[0].operator = 0x1374038C2E2403f9aB7db62EE7516e0119F1124A;
        operatorKickParams[1].quorumNumber = 0x1;
        operatorKickParams[1].operator = 0xD393FD495367164d7eB53840e59469c13266bA59;

        // Values are random
        ISignatureUtilsMixinTypes.SignatureWithSaltAndExpiry memory churnApproverSignature;
        churnApproverSignature.signature =
            hex"d547fa0126f97d1752a3b3103c495961a4a6a7a5386feb32ac514289c578db5a0d64bfa34855c39d78cce241a9c73d5cae47dee02c691fd5320fdef4ad3e1f8e";
        churnApproverSignature.salt = hex"7879ea091cd16d7afec6bc1e96b92f2229f744c703fb9603b2ca6f60ea9df6c0";
        churnApproverSignature.expiry = 138752197623537982159531315300136159918886501617089726422768086017712946835;

        assertEq(
            abi.encode(registrationType, socket, params, operatorKickParams, churnApproverSignature),
            hex"0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000018001ac6045296d64b31ed644e53ce1a1c4f72f67a2d47b06b652ca8167f1b2ada900fb7cd59f322f4dffa18360bfcdc15f1f6cd09aea573efa919dc828dfabaf5f16ee091592629fc566636de7b3d53322f4833014b4655dd57279cfb2828bbc6e0c5cb58c8d572dc9dcf5f5999501533e22243fac4f8ee4452a6dd0d4bcaeb2e403006a43453d56eafa7dc4ddcbd41b2330031f58e437ee3806c50a9e554a0cbc0bae792831463d56a1a9983647b77fcdbae38a6621ceef9e147614bb4869bf36267e47def74c144f8e7238dd088097943d1007f8ce7cab028151e9a1beac6d500e56fef5ea67a586d1fbb03dcc0a268f6c4835ad1c215eadc77cc676c378101d00000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002600000000000000000000000000000000000000000000000000000000000000006756e757365640000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000001374038c2e2403f9ab7db62ee7516e0119f1124a0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000d393fd495367164d7eb53840e59469c13266ba5900000000000000000000000000000000000000000000000000000000000000607879ea091cd16d7afec6bc1e96b92f2229f744c703fb9603b2ca6f60ea9df6c0004e87ed0c684886b5b2e661e58348383df270f2ad5630aaadde83c88c517e930000000000000000000000000000000000000000000000000000000000000040d547fa0126f97d1752a3b3103c495961a4a6a7a5386feb32ac514289c578db5a0d64bfa34855c39d78cce241a9c73d5cae47dee02c691fd5320fdef4ad3e1f8e"
        );
    }
}
