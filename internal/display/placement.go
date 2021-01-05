package display

func Scaffold() {
	_, dispY := Screen.Size()
	PrintLine(dispY - BottomLineHeight, '-', false)

	PrintString(">", 0, dispY - BottomLineHeight + 1)

}

func FindTopLeftCoord(imgX, imgY, sizeX, sizeY int) (int, int) {
	remX := sizeX - imgX
	remY := sizeY - imgY
	return remX / 2, remY / 2
}