//go:build (sdk_net_proto || sdk || all) && !no_sdk && !no_sdk_net_proto

package modules

import (
	_ "github.com/ZenLiuCN/engineX/modules/golang/net/mail"
	_ "github.com/ZenLiuCN/engineX/modules/golang/net/netip"
	_ "github.com/ZenLiuCN/engineX/modules/golang/net/smtp"
	_ "github.com/ZenLiuCN/engineX/modules/golang/net/textproto"
)
