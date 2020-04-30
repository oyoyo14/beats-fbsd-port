+++ vendor/github.com/insomniacslk/dhcp/dhcpv4/bindtodevice_bsd.go     2020-04-30 08:47:05.639979000 +0000
@@ -0,0 +1,18 @@
+// +build freebsd openbsd netbsd
+-
+package dhcpv4
+-
+import (
+      "net"
+      "syscall"
+)
+-
+// BindToInterface emulates linux's SO_BINDTODEVICE option for a socket by using
+// IP_RECVIF.
+func BindToInterface(fd int, ifname string) error {
+      iface, err := net.InterfaceByName(ifname)
+      if err != nil {
+              return err
+      }
+      return syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_RECVIF, iface.Index)
+}
