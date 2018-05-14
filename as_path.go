package main
//
//import "fmt"
//
//type PathManager string
//
//func (pm PathManager) path(name string, params string) string {
//	return fmt.Sprintf("java -jar %s%s %s", pm, name, params)
//}
//
//func (pm PathManager) clientParams(hostname string, port int, username, secret string) string {
//	return fmt.Sprintf("-rh %s -rp %d -u %s -s %s", hostname, port, username, secret)
//}
//
//func (pm PathManager) serverParams(hostname, remoteHostname string, port, remotePort int, secret string) string {
//	return fmt.Sprintf("-h %s -p %d -rh %s -rp %d -s %s", hostname, port, remoteHostname, remotePort, secret)
//}
//
//func (pm PathManager) Dir() string {
//	return string(pm)
//}
//
//func (pm PathManager) Client(hostname string, port int, username, secret string) string {
//	return pm.path("Client.jar", pm.clientParams(hostname, port, username, secret))
//}
//
//func (pm PathManager) Master(hostname string, port int, secret string) string {
//	return pm.path("Master.jar", pm.serverParams(hostname, "", port, 0, secret))
//}
//func (pm PathManager) Slave(hostname string, port int, remoteHostname string, remotePort int, secret string) string {
//	return pm.path("Slave.jar", pm.serverParams(hostname, remoteHostname, port, remotePort, secret))
//}
//func (pm PathManager) Gateway(hostname string, port int, remoteHostname string, remotePort int, secret string) string {
//	return pm.path("Gateway.jar", pm.serverParams(hostname, remoteHostname, port, remotePort, secret))
//}
//
//func NewPathManager() PathManager {
//	return PathManager("/Users/shixiao/Playground/COMP90015/out/artifacts/jar/")
//}
