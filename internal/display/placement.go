package display

const (
	infoBoxSizeX = 35
	infoBoxSizeY = 7
)

func Scaffold() {
	_, dispY := Screen.Size()
	Box(2, 1, 2+infoBoxSizeX, 1+infoBoxSizeY, " STATS ")
	PrintLine(dispY - BottomLineHeight, '─', false)
	PrintString(">", 0, dispY - BottomLineHeight + 2)

}

func FindTopLeftCoord(imgX, imgY, sizeX, sizeY int) (int, int) {
	remX := sizeX - imgX
	remY := sizeY - imgY
	return remX / 2, remY / 2
}