package model


// 代理类型
const (
    PXY_TYPE_UNKNOWN byte = iota
    PXY_TYPE_HTTP
    PXY_TYPE_HTTPS
    PXY_TYPE_HTTP_SOCKS4
    PXY_TYPE_HTTP_SOCKS5
)

// 匿名程度
const (
    //未知类型
    ANONYMITY_LEVEL_UNKNOWN byte = iota
    // 透明代理
    ANONYMITY_LEVEL_OPEN
    // 普通代理
    ANONYMITY_LEVEL_GENERAL
    // 高匿代理
    ANONYMITY_LEVEL_ANVANCED
)

/* 获取代理ip 信息
{
    ip地址
    端口
    代理类型
    匿名程度
}
*/
type IpInfo struct {
    Ip      string
    Port    string
    PxyType byte
    Level   byte
}
