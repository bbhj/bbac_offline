<?php
$appid = "200001";
$bucket = "newbucket";
$secret_id = "AKIDUfLUEUigQiXqm7CVSspKJnuaiIKtxqAv";
$secret_key = "bLcPnl88WU30VY57ipRhSePfPdOfSruK";
$expired = 1437995644 + 60;
$onceExpired = 0;
$current = 1437995644;
$rdm = 2081660421;
$fileid = "/200001/newbucket/tencent_test.jpg";

$multi_effect_signature = 'a='.$appid.'&b='.$bucket.'&k='.$secret_id.'&e='.$expired.'&t='.$current.'&r='.$rdm.'&f=';
echo $multi_effect_signature."\n";

$multi_effect_signature = base64_encode(hash_hmac('SHA1', $multi_effect_signature, $secret_key, true).$multi_effect_signature);
echo $multi_effect_signature."\n"

?>
