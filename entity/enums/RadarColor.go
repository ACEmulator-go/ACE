package enums

type RadarColor byte

const (
	RadarColorDefault          RadarColor = RadarColor(0x00)
	RadarColorBlue             RadarColor = RadarColor(0x01)
	RadarColorGold             RadarColor = RadarColor(0x02)
	RadarColorWhite            RadarColor = RadarColor(0x03)
	RadarColorPurple           RadarColor = RadarColor(0x04)
	RadarColorRed              RadarColor = RadarColor(0x05)
	RadarColorPink             RadarColor = RadarColor(0x06)
	RadarColorGreen            RadarColor = RadarColor(0x07)
	RadarColorYellow           RadarColor = RadarColor(0x08)
	RadarColorCyan             RadarColor = RadarColor(0x09)
	RadarColorBrightGreen      RadarColor = RadarColor(0x10)
	RadarColorAdmin            RadarColor = RadarColor(RadarColorCyan)
	RadarColorAdvocate         RadarColor = RadarColor(RadarColorPink)
	RadarColorCreature         RadarColor = RadarColor(RadarColorGold)
	RadarColorLifeStone        RadarColor = RadarColor(RadarColorBlue)
	RadarColorNPC              RadarColor = RadarColor(RadarColorYellow)
	RadarColorPlayerKiller     RadarColor = RadarColor(RadarColorRed)
	RadarColorPortal           RadarColor = RadarColor(RadarColorPurple)
	RadarColorSentinel         RadarColor = RadarColor(RadarColorCyan)
	RadarColorVendor           RadarColor = RadarColor(RadarColorYellow)
	RadarColorFellowship       RadarColor = RadarColor(RadarColorBrightGreen)
	RadarColorFellowshipLeader RadarColor = RadarColor(RadarColorBrightGreen)
)
