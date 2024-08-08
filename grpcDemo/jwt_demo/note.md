# JWT
JWT（JSON Web Token）是一种开放标准（RFC 7519），用于在网络应用间安全地传输声明。它可以通过数字签名或加密保证声明的完整性，并且通常会被用来在用户和服务器之间传递身份信息。在 Golang 中，通过使用第三方库可以轻松地实现 JWT 的生成、签名和验证功能。本文将介绍如何使用 Golang 实现 JWT 的基本功能，并提供了一个简单的示例代码


## JWT构成
由三个部分组成的字符串，它们以点分隔开：

    Header（头部）：
        包含了 Token 的元数据，例如类型（JWT）和所使用的算法（例如 HMAC SHA256 或 RSA）。
    Payload（载荷）：
        包含了声明（claim），声明是关于实体（通常是用户）和其他数据的一些声明性的信息。
        声明可以分为三种类型：
            注册声明
                这些声明是预定义的，非必须使用的但被推荐使用。官方标准定义的注册声明有 7 个：
                    iss(Issuer)	发行者，标识 JWT 的发行者。
                    sub(Subject)	主题，标识 JWT 的主题，通常指用户的唯一标识
                    aud(Audience)	观众，标识 JWT的接收者
                    exp(Expiration Time)	过期时间。标识 JWT 的过期时间，这个时间必须是将来的
                    nbf(Not Before)	不可用时间。在此时间之前，JWT 不应被接受处理
                    iat(Issued At)	发行时间，标识 JWT 的发行时间
                    jti(JWT ID)	JWT 的唯一标识符，用于防止 JWT 被重放（即重复使用）
            公共声明
                可以由使用 JWT 的人自定义，但为了避免冲突，任何新定义的声明都应已在 IANA JSON Web Token Registry 中注册或者是一个 公共名称，其中包含了碰撞防抗性名称（Collision-Resistant Name）。
            私有声明
                发行和使用 JWT 的双方共同商定的声明，区别于 注册声明 和 公共声明。

    Signature（签名）：
        为了防止数据篡改，将头部Header和负载Payload的信息进行一定算法处理，加上一个密钥，最后生成签名。

