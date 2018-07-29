package QemuCmd

import (
	"testing"
	"log"
)

/*
func TestMipsArgs_LoadImage(t *testing.T) {
	image_str := "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/file/debian_squeeze_mips_lanimei123.qcow2"
	mount_str := "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/test"
	err := LoadImage(image_str, mount_str)
	if err != nil {
		log.Fatal(err)
	}
}


func TestMipsArgs_UnloadImage(t *testing.T) {
	mount_str := "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/test"
	err := UnloadImage(mount_str)
	if err != nil {
		log.Println(err)
	}
}

*/
/*
func TestMakedir(t *testing.T) {
	err := Makedir("/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/test")
	if err != nil {
		log.Fatal(err)
	}
}
*/

/*
func TestCpInitImage(t *testing.T) {
	err := CpInitImage("/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/file/rc.local", "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/")
	if err != nil {
		log.Fatal(err)
	}
}
*/

/*
func TestCreateImage(t *testing.T) {
	filesPath := FilesPath{
		MalwarePath: "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/malware_one",
		StartScript: "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/file/rc.local",
		ImageSrcPath:"/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/debian_squeeze_mips_standard.qcow2",
		ImageNameSpecify: "debian_squeeze_mips_lanimei123.qcow2",
		ImageDestPath: "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/file",
		LoadPath: "/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/test",
	}
	ImageErr := CreateImage(filesPath)
	if ImageErr != nil {
		log.Fatal(ImageErr)
	}
}
*/

/*
func TestMalwareClassify(t *testing.T) {
	output, err := MalwareClassify("/home/lanimei/lanimei_work/Bread_mal/malware_07_09/mal/MIPS/ff1e7def7f1c15ae17b66256eda2ae26b35841a589860953389a7cc25ca1fef5")
	if err != nil {
		log.Fatal("文件分类出现问题")
	}
	log.Println(output)
	output, err = MalwareClassify("/home/lanimei/lanimei_work/Bread_mal/malware_07_09/mal/RenesasSH/fe92a75a92064a5effc3ab8ebdb80e08c9065ac2ec0af8fb3fb0a475b425b4a7")
	if err != nil {
		log.Fatal("文件分类出现问题")
	}
	log.Println(output)
	output, err = MalwareClassify("/home/lanimei/lanimei_work/Bread_mal/malware_07_09/mal/ARM/ffad24af27160c5146aedf56f073a1b5de7f03699bd9828da6fdb6ac5316ece5")
	if err != nil {
		log.Fatal("文件分类出现问题")
	}
	log.Println(output)
	output, err = MalwareClassify("/home/lanimei/lanimei_work/Bread_mal/malware_07_09/mal/80386/fec01bdab80b7b50174be006c8e85ddc4ac37489bca4e18cb3d75e3fd62dd702")
	if err != nil {
		log.Fatal("文件分类出现问题")
	}
	log.Println(output)
	output, err = MalwareClassify("/home/lanimei/lanimei_work/Bread_mal/malware_07_09/mal/x86-64/f97729dc16252004cdd4fc144d55955b75c369f91f6ecc2e01b40fef1374fd26")
	if err != nil {
		log.Fatal("文件分类出现问题")
	}
	log.Println(output)
	output, err = MalwareClassify("/home/lanimei/lanimei_work/Bread_mal/malware_07_09/mal/SPARC/ffbd4e4bd8464e8f724a3a0cfa6a660b95869af09f6c74b6c9b7d8fea419d501")
	if err != nil {
		log.Fatal("文件分类出现问题")
	}
	log.Println(output)
	output, err = MalwareClassify("/home/lanimei/lanimei_work/Bread_mal/malware_07_09/mal/PowerPC/fee93690d50a596b98eb7f1990a2cc11e3969fecbbd964a9ca3ce1c0ac48b4ad")
	if err != nil {
		log.Fatal("文件分类出现问题")
	}
	log.Println(output)
}

func TestGetAllMalware(t *testing.T) {
	Malwarelist, err := GetAllMalware("/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/malware")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(Malwarelist)
}

func TestGetMalwareSha256(t *testing.T) {
	Sha256, err := GetMalwareSha256("/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/rc.local")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(Sha256)
}

func TestParseConfig(t *testing.T) {
	lanimei, err := ParseConfig("/home/lanimei/Project/go/src/capture_packet/CaptureProject/QemuCmd/test.json")
	if err != nil {
		log.Fatal("解析文件出现问题")
		log.Println(err)
	}
	log.Println("MalwarePath: ", lanimei.MalwarePath)
	log.Println("StartScript: ", lanimei.StartScript)
	log.Println("ImageSrcPath: ", lanimei.ImageSrcPath)
	log.Println("ImageDestPath: ", lanimei.ImageDestPath)
	log.Println("LoadPath: ", lanimei.LoadPath)
}
*/

func TestInitFilesPath(t *testing.T) {
	log.Println("开始：")
	InitFilesPath()   //主函数中执行该函数进行赋值
	log.Println(len(Malwares))
	for _, item := range Malwares{
		item.CreateStartup()
		item.CreateImage()
	}
}