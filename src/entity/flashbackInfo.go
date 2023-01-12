package entity

import "strings"

type FlashbackInfo struct {
	sourceUserInfo string
	sourceSocket   string
	sourceIp       string
	sourcePort     string
	targetUserInfo string
	targetSocket   string
	targetIp       string
	targetPort     string
	sshUser        string
	sshPass        string
	sshInfo        string
	sshPort        string
}

func (f *FlashbackInfo) SourceIp() string {
	return f.sourceIp
}

func (f *FlashbackInfo) SetSourceIp(sourceIp string) {
	f.sourceIp = sourceIp
}

func (f *FlashbackInfo) SourcePort() string {
	return f.sourcePort
}

func (f *FlashbackInfo) SetSourcePort(sourcePort string) {
	f.sourcePort = sourcePort
}

func (f *FlashbackInfo) TargetIp() string {
	return f.targetIp
}

func (f *FlashbackInfo) SetTargetIp(targetIp string) {
	f.targetIp = targetIp
}

func (f *FlashbackInfo) TargetPort() string {
	return f.targetPort
}

func (f *FlashbackInfo) SetTargetPort(targetPort string) {
	f.targetPort = targetPort
}

func (f *FlashbackInfo) SshPort() string {
	return f.sshPort
}

func (f *FlashbackInfo) SetSshPort(sshPort string) {
	f.sshPort = sshPort
}

func (f *FlashbackInfo) SourceUserInfo() string {
	return f.sourceUserInfo
}

func (f *FlashbackInfo) SetSourceUserInfo(sourceUserInfo string) {
	f.sourceUserInfo = sourceUserInfo
}

func (f *FlashbackInfo) SourceSocket() string {
	return f.sourceSocket
}

func (f *FlashbackInfo) SetSourceSocket(sourceSocket string) {
	f.sourceSocket = sourceSocket
	f.SetSourceIp(strings.Split(sourceSocket, ":")[0])
	f.SetSourcePort(strings.Split(sourceSocket, ":")[1])
}

func (f *FlashbackInfo) TargetUserInfo() string {
	return f.targetUserInfo
}

func (f *FlashbackInfo) SetTargetUserInfo(targetUserInfo string) {
	f.targetUserInfo = targetUserInfo
}

func (f *FlashbackInfo) TargetSocket() string {
	return f.targetSocket
}

func (f *FlashbackInfo) SetTargetSocket(targetSocket string) {
	f.targetSocket = targetSocket
	f.SetTargetIp(strings.Split(targetSocket, ":")[0])
	f.SetTargetPort(strings.Split(targetSocket, ":")[1])
}

func (f *FlashbackInfo) SshUser() string {
	return f.sshUser
}

func (f *FlashbackInfo) SetSshUser(sshUser string) {
	f.sshUser = sshUser
}

func (f *FlashbackInfo) SshPass() string {
	return f.sshPass
}

func (f *FlashbackInfo) SetSshPass(sshPass string) {
	f.sshPass = sshPass
}

func (f *FlashbackInfo) SshInfo() string {
	return f.sshInfo
}

func (f *FlashbackInfo) SetSshInfo(sshInfo string) {
	f.sshInfo = sshInfo
}
