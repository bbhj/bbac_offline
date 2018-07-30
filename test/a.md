$multi_effect_signature = base64_encode(hash_hmac('SHA1', $multi_effect_signature, $secret_key, true).$multi_effect_signature);

在php中hash_hmac函数就能将HMAC和一部分哈希加密算法相结合起来实现HMAC-SHA1  HMAC-SHA256 HMAC-MD5等等算法。函数介绍如下：
string hash_hmac(string $algo, string $data, string $key, bool $raw_output = false)
algo：要使用的哈希算法名称，可以是上述提到的md5,sha1等
data：要进行哈希运算的消息，也就是需要加密的明文。
key：使用HMAC生成信息摘要是所使用的密钥。
raw_output：该参数为可选参数，默认为false，如果设为true，则返回原始二进制数据表示的信息摘要，否则返回16进制小写字符串格式表示的信息摘要（注意是16进制数，而非简单的字母加数字）。
另外：如果algo参数指定的不是受支持的算法，将返回false。
