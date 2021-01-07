package display

const (
	infoBoxSizeX = 35
	infoBoxSizeY = 7

	infoBoxPosX = 2
	infoBoxPosY = 1
)

func Scaffold() {
	_, dispY := Screen.Size()
	Box(infoBoxPosX, infoBoxPosY, infoBoxPosX+infoBoxSizeX, infoBoxPosY+infoBoxSizeY, " STATS ")
	PrintLine(dispY - BottomLineHeight, '─', false)
	PrintString(">", 0, dispY - BottomLineHeight + 2)

}

func FindTopLeftCoord(imgX, imgY, sizeX, sizeY int) (int, int) {
	remX := sizeX - imgX
	remY := sizeY - imgY
	return remX / 2, remY / 2
}