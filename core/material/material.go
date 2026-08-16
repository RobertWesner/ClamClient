package material

type Material interface {
	IsMaterial()
}

type Base struct {
	ID        int
	StackSize int
}

func (m Base) IsMaterial() {}

type WithData[T Data] struct {
	Base
}

func (wd WithData[T]) IsMaterial() {}
func (wd WithData[T]) Data(raw uint8) T {
	return T(raw)
}

type WithDurability struct {
	Base
	Durability int
}

func (wd WithDurability) IsMaterial() {}

var (
	Air                 = Base{0, 64}
	Stone               = Base{1, 64}
	Grass               = Base{2, 64}
	Dirt                = Base{3, 64}
	Cobblestone         = Base{4, 64}
	Wood                = Base{5, 64}
	Sapling             = WithData[DataTree]{Base{6, 64}}
	Bedrock             = Base{7, 64}
	Water               = WithData[DataLiquid]{Base{8, 64}}
	StationaryWater     = WithData[DataLiquid]{Base{9, 64}}
	Lava                = WithData[DataLiquid]{Base{10, 64}}
	StationaryLava      = WithData[DataLiquid]{Base{11, 64}}
	Sand                = Base{12, 64}
	Gravel              = Base{13, 64}
	GoldOre             = Base{14, 64}
	IronOre             = Base{15, 64}
	CoalOre             = Base{16, 64}
	Log                 = Base{17, 64}
	Leaves              = Base{18, 64}
	Sponge              = Base{19, 64}
	Glass               = Base{20, 64}
	LapisOre            = Base{21, 64}
	LapisBlock          = Base{22, 64}
	Dispenser           = WithData[DataDispenser]{Base{23, 64}}
	Sandstone           = Base{24, 64}
	NoteBlock           = Base{25, 64}
	BedBlock            = WithData[DataBed]{Base{26, 64}}
	PoweredRail         = WithData[DataPoweredRail]{Base{27, 64}}
	DetectorRail        = WithData[DataDetectorRail]{Base{28, 64}}
	PistonStickyBase    = WithData[DataPistonBaseMaterial]{Base{29, 64}}
	Web                 = Base{30, 64}
	LongGrass           = WithData[DataLongGrass]{Base{31, 64}}
	DeadBush            = Base{32, 64}
	PistonBase          = WithData[DataPistonBaseMaterial]{Base{33, 64}}
	PistonExtension     = WithData[DataPistonExtensionMaterial]{Base{34, 64}}
	Wool                = WithData[DataWool]{Base{35, 64}}
	PistonMovingPiece   = Base{36, 64}
	YellowFlower        = Base{37, 64}
	RedRose             = Base{38, 64}
	BrownMushroom       = Base{39, 64}
	RedMushroom         = Base{40, 64}
	GoldBlock           = Base{41, 64}
	IronBlock           = Base{42, 64}
	DoubleStep          = WithData[DataStep]{Base{43, 64}}
	Step                = WithData[DataStep]{Base{44, 64}}
	Brick               = Base{45, 64}
	Tnt                 = Base{46, 64}
	Bookshelf           = Base{47, 64}
	MossyCobblestone    = Base{48, 64}
	Obsidian            = Base{49, 64}
	Torch               = WithData[DataTorch]{Base{50, 64}}
	Fire                = Base{51, 64}
	MobSpawner          = Base{52, 64}
	WoodStairs          = WithData[DataStairs]{Base{53, 64}}
	Chest               = Base{54, 64}
	RedstoneWire        = WithData[DataRedstoneWire]{Base{55, 64}}
	DiamondOre          = Base{56, 64}
	DiamondBlock        = Base{57, 64}
	Workbench           = Base{58, 64}
	Crops               = WithData[DataCrops]{Base{59, 64}}
	Soil                = WithData[DataSoil]{Base{60, 64}}
	Furnace             = WithData[DataFurnace]{Base{61, 1}}
	BurningFurnace      = WithData[DataFurnace]{Base{62, 64}}
	SignPost            = WithData[DataSign]{Base{63, 1}}
	WoodenDoor          = WithData[DataDoor]{Base{64, 64}}
	Ladder              = WithData[DataLadder]{Base{65, 64}}
	Rails               = WithData[DataRails]{Base{66, 64}}
	CobblestoneStairs   = WithData[DataStairs]{Base{67, 64}}
	WallSign            = WithData[DataSign]{Base{68, 1}}
	Lever               = WithData[DataLever]{Base{69, 64}}
	StonePlate          = WithData[DataPressurePlate]{Base{70, 64}}
	IronDoorBlock       = WithData[DataDoor]{Base{71, 64}}
	WoodPlate           = WithData[DataPressurePlate]{Base{72, 64}}
	RedstoneOre         = Base{73, 64}
	GlowingRedstoneOre  = Base{74, 64}
	RedstoneTorchOff    = WithData[DataRedstoneTorch]{Base{75, 64}}
	RedstoneTorchOn     = WithData[DataRedstoneTorch]{Base{76, 64}}
	StoneButton         = WithData[DataButton]{Base{77, 64}}
	Snow                = Base{78, 64}
	Ice                 = Base{79, 64}
	SnowBlock           = Base{80, 64}
	Cactus              = WithData[DataCactus]{Base{81, 64}}
	Clay                = Base{82, 64}
	SugarCaneBlock      = WithData[DataSugarCaneBlock]{Base{83, 64}}
	Jukebox             = WithData[DataJukebox]{Base{84, 64}}
	Fence               = Base{85, 64}
	Pumpkin             = WithData[DataPumpkin]{Base{86, 64}}
	Netherrack          = Base{87, 64}
	SoulSand            = Base{88, 64}
	Glowstone           = Base{89, 64}
	Portal              = Base{90, 64}
	JackOLantern        = WithData[DataPumpkin]{Base{91, 64}}
	CakeBlock           = WithData[DataCake]{Base{92, 1}}
	DiodeBlockOff       = WithData[DataDiode]{Base{93, 64}}
	DiodeBlockOn        = WithData[DataDiode]{Base{94, 64}}
	LockedChest         = Base{95, 64}
	TrapDoor            = WithData[DataTrapdoor]{Base{96, 64}}
	IronSpade           = WithDurability{Base{256, 1}, 250}
	IronPickaxe         = WithDurability{Base{257, 1}, 250}
	IronAxe             = WithDurability{Base{258, 1}, 250}
	FlintAndSteel       = WithDurability{Base{259, 1}, 64}
	Apple               = Base{260, 1}
	Bow                 = Base{261, 1}
	Arrow               = Base{262, 64}
	Coal                = WithData[DataCoal]{Base{263, 64}}
	Diamond             = Base{264, 64}
	IronIngot           = Base{265, 64}
	GoldIngot           = Base{266, 64}
	IronSword           = WithDurability{Base{267, 1}, 250}
	WoodSword           = WithDurability{Base{268, 1}, 59}
	WoodSpade           = WithDurability{Base{269, 1}, 59}
	WoodPickaxe         = WithDurability{Base{270, 1}, 59}
	WoodAxe             = WithDurability{Base{271, 1}, 59}
	StoneSword          = WithDurability{Base{272, 1}, 131}
	StoneSpade          = WithDurability{Base{273, 1}, 131}
	StonePickaxe        = WithDurability{Base{274, 1}, 131}
	StoneAxe            = WithDurability{Base{275, 1}, 131}
	DiamondSword        = WithDurability{Base{276, 1}, 1561}
	DiamondSpade        = WithDurability{Base{277, 1}, 1561}
	DiamondPickaxe      = WithDurability{Base{278, 1}, 1561}
	DiamondAxe          = WithDurability{Base{279, 1}, 1561}
	Stick               = Base{280, 64}
	Bowl                = Base{281, 64}
	MushroomSoup        = Base{282, 1}
	GoldSword           = WithDurability{Base{283, 1}, 32}
	GoldSpade           = WithDurability{Base{284, 1}, 32}
	GoldPickaxe         = WithDurability{Base{285, 1}, 32}
	GoldAxe             = WithDurability{Base{286, 1}, 32}
	String              = Base{287, 64}
	Feather             = Base{288, 64}
	Sulphur             = Base{289, 64}
	WoodHoe             = WithDurability{Base{290, 1}, 59}
	StoneHoe            = WithDurability{Base{291, 1}, 131}
	IronHoe             = WithDurability{Base{292, 1}, 250}
	DiamondHoe          = WithDurability{Base{293, 1}, 1561}
	GoldHoe             = WithDurability{Base{294, 1}, 32}
	Seeds               = Base{295, 64}
	Wheat               = Base{296, 64}
	Bread               = Base{297, 1}
	LeatherHelmet       = WithDurability{Base{298, 1}, 33}
	LeatherChestplate   = WithDurability{Base{299, 1}, 47}
	LeatherLeggings     = WithDurability{Base{300, 1}, 45}
	LeatherBoots        = WithDurability{Base{301, 1}, 39}
	ChainmailHelmet     = WithDurability{Base{302, 1}, 66}
	ChainmailChestplate = WithDurability{Base{303, 1}, 95}
	ChainmailLeggings   = WithDurability{Base{304, 1}, 91}
	ChainmailBoots      = WithDurability{Base{305, 1}, 78}
	IronHelmet          = WithDurability{Base{306, 1}, 135}
	IronChestplate      = WithDurability{Base{307, 1}, 191}
	IronLeggings        = WithDurability{Base{308, 1}, 183}
	IronBoots           = WithDurability{Base{309, 1}, 159}
	DiamondHelmet       = WithDurability{Base{310, 1}, 271}
	DiamondChestplate   = WithDurability{Base{311, 1}, 383}
	DiamondLeggings     = WithDurability{Base{312, 1}, 367}
	DiamondBoots        = WithDurability{Base{313, 1}, 319}
	GoldHelmet          = WithDurability{Base{314, 1}, 67}
	GoldChestplate      = WithDurability{Base{315, 1}, 95}
	GoldLeggings        = WithDurability{Base{316, 1}, 91}
	GoldBoots           = WithDurability{Base{317, 1}, 79}
	Flint               = Base{318, 64}
	Pork                = Base{319, 1}
	GrilledPork         = Base{320, 1}
	Painting            = Base{321, 64}
	GoldenApple         = Base{322, 1}
	Sign                = Base{323, 1}
	WoodDoor            = Base{324, 1}
	Bucket              = Base{325, 1}
	WaterBucket         = Base{326, 1}
	LavaBucket          = Base{327, 1}
	Minecart            = Base{328, 1}
	Saddle              = Base{329, 1}
	IronDoor            = Base{330, 1}
	Redstone            = Base{331, 64}
	SnowBall            = Base{332, 16}
	Boat                = Base{333, 1}
	Leather             = Base{334, 64}
	MilkBucket          = Base{335, 1}
	ClayBrick           = Base{336, 64}
	ClayBall            = Base{337, 64}
	SugarCane           = Base{338, 64}
	Paper               = Base{339, 64}
	Book                = Base{340, 64}
	SlimeBall           = Base{341, 64}
	StorageMinecart     = Base{342, 1}
	PoweredMinecart     = Base{343, 1}
	Egg                 = Base{344, 16}
	Compass             = Base{345, 64}
	FishingRod          = WithDurability{Base{346, 1}, 64}
	Watch               = Base{347, 64}
	GlowstoneDust       = Base{348, 64}
	RawFish             = Base{349, 1}
	CookedFish          = Base{350, 1}
	InkSack             = WithData[DataDye]{Base{351, 64}}
	Bone                = Base{352, 64}
	Sugar               = Base{353, 64}
	Cake                = Base{354, 1}
	Bed                 = Base{355, 1}
	Diode               = Base{356, 64}
	Cookie              = Base{357, 8}
	Map                 = WithData[DataMap]{Base{358, 1}}
	Shears              = WithDurability{Base{359, 1}, 238}
	GoldRecord          = Base{2256, 1}
	GreenRecord         = Base{2257, 1}
)

func FromID(id int) Material {
	switch id {
	case 0:
		return Air
	case 1:
		return Stone
	case 2:
		return Grass
	case 3:
		return Dirt
	case 4:
		return Cobblestone
	case 5:
		return Wood
	case 6:
		return Sapling
	case 7:
		return Bedrock
	case 8:
		return Water
	case 9:
		return StationaryWater
	case 10:
		return Lava
	case 11:
		return StationaryLava
	case 12:
		return Sand
	case 13:
		return Gravel
	case 14:
		return GoldOre
	case 15:
		return IronOre
	case 16:
		return CoalOre
	case 17:
		return Log
	case 18:
		return Leaves
	case 19:
		return Sponge
	case 20:
		return Glass
	case 21:
		return LapisOre
	case 22:
		return LapisBlock
	case 23:
		return Dispenser
	case 24:
		return Sandstone
	case 25:
		return NoteBlock
	case 26:
		return BedBlock
	case 27:
		return PoweredRail
	case 28:
		return DetectorRail
	case 29:
		return PistonStickyBase
	case 30:
		return Web
	case 31:
		return LongGrass
	case 32:
		return DeadBush
	case 33:
		return PistonBase
	case 34:
		return PistonExtension
	case 35:
		return Wool
	case 36:
		return PistonMovingPiece
	case 37:
		return YellowFlower
	case 38:
		return RedRose
	case 39:
		return BrownMushroom
	case 40:
		return RedMushroom
	case 41:
		return GoldBlock
	case 42:
		return IronBlock
	case 43:
		return DoubleStep
	case 44:
		return Step
	case 45:
		return Brick
	case 46:
		return Tnt
	case 47:
		return Bookshelf
	case 48:
		return MossyCobblestone
	case 49:
		return Obsidian
	case 50:
		return Torch
	case 51:
		return Fire
	case 52:
		return MobSpawner
	case 53:
		return WoodStairs
	case 54:
		return Chest
	case 55:
		return RedstoneWire
	case 56:
		return DiamondOre
	case 57:
		return DiamondBlock
	case 58:
		return Workbench
	case 59:
		return Crops
	case 60:
		return Soil
	case 61:
		return Furnace
	case 62:
		return BurningFurnace
	case 63:
		return SignPost
	case 64:
		return WoodenDoor
	case 65:
		return Ladder
	case 66:
		return Rails
	case 67:
		return CobblestoneStairs
	case 68:
		return WallSign
	case 69:
		return Lever
	case 70:
		return StonePlate
	case 71:
		return IronDoorBlock
	case 72:
		return WoodPlate
	case 73:
		return RedstoneOre
	case 74:
		return GlowingRedstoneOre
	case 75:
		return RedstoneTorchOff
	case 76:
		return RedstoneTorchOn
	case 77:
		return StoneButton
	case 78:
		return Snow
	case 79:
		return Ice
	case 80:
		return SnowBlock
	case 81:
		return Cactus
	case 82:
		return Clay
	case 83:
		return SugarCaneBlock
	case 84:
		return Jukebox
	case 85:
		return Fence
	case 86:
		return Pumpkin
	case 87:
		return Netherrack
	case 88:
		return SoulSand
	case 89:
		return Glowstone
	case 90:
		return Portal
	case 91:
		return JackOLantern
	case 92:
		return CakeBlock
	case 93:
		return DiodeBlockOff
	case 94:
		return DiodeBlockOn
	case 95:
		return LockedChest
	case 96:
		return TrapDoor
	case 256:
		return IronSpade
	case 257:
		return IronPickaxe
	case 258:
		return IronAxe
	case 259:
		return FlintAndSteel
	case 260:
		return Apple
	case 261:
		return Bow
	case 262:
		return Arrow
	case 263:
		return Coal
	case 264:
		return Diamond
	case 265:
		return IronIngot
	case 266:
		return GoldIngot
	case 267:
		return IronSword
	case 268:
		return WoodSword
	case 269:
		return WoodSpade
	case 270:
		return WoodPickaxe
	case 271:
		return WoodAxe
	case 272:
		return StoneSword
	case 273:
		return StoneSpade
	case 274:
		return StonePickaxe
	case 275:
		return StoneAxe
	case 276:
		return DiamondSword
	case 277:
		return DiamondSpade
	case 278:
		return DiamondPickaxe
	case 279:
		return DiamondAxe
	case 280:
		return Stick
	case 281:
		return Bowl
	case 282:
		return MushroomSoup
	case 283:
		return GoldSword
	case 284:
		return GoldSpade
	case 285:
		return GoldPickaxe
	case 286:
		return GoldAxe
	case 287:
		return String
	case 288:
		return Feather
	case 289:
		return Sulphur
	case 290:
		return WoodHoe
	case 291:
		return StoneHoe
	case 292:
		return IronHoe
	case 293:
		return DiamondHoe
	case 294:
		return GoldHoe
	case 295:
		return Seeds
	case 296:
		return Wheat
	case 297:
		return Bread
	case 298:
		return LeatherHelmet
	case 299:
		return LeatherChestplate
	case 300:
		return LeatherLeggings
	case 301:
		return LeatherBoots
	case 302:
		return ChainmailHelmet
	case 303:
		return ChainmailChestplate
	case 304:
		return ChainmailLeggings
	case 305:
		return ChainmailBoots
	case 306:
		return IronHelmet
	case 307:
		return IronChestplate
	case 308:
		return IronLeggings
	case 309:
		return IronBoots
	case 310:
		return DiamondHelmet
	case 311:
		return DiamondChestplate
	case 312:
		return DiamondLeggings
	case 313:
		return DiamondBoots
	case 314:
		return GoldHelmet
	case 315:
		return GoldChestplate
	case 316:
		return GoldLeggings
	case 317:
		return GoldBoots
	case 318:
		return Flint
	case 319:
		return Pork
	case 320:
		return GrilledPork
	case 321:
		return Painting
	case 322:
		return GoldenApple
	case 323:
		return Sign
	case 324:
		return WoodDoor
	case 325:
		return Bucket
	case 326:
		return WaterBucket
	case 327:
		return LavaBucket
	case 328:
		return Minecart
	case 329:
		return Saddle
	case 330:
		return IronDoor
	case 331:
		return Redstone
	case 332:
		return SnowBall
	case 333:
		return Boat
	case 334:
		return Leather
	case 335:
		return MilkBucket
	case 336:
		return ClayBrick
	case 337:
		return ClayBall
	case 338:
		return SugarCane
	case 339:
		return Paper
	case 340:
		return Book
	case 341:
		return SlimeBall
	case 342:
		return StorageMinecart
	case 343:
		return PoweredMinecart
	case 344:
		return Egg
	case 345:
		return Compass
	case 346:
		return FishingRod
	case 347:
		return Watch
	case 348:
		return GlowstoneDust
	case 349:
		return RawFish
	case 350:
		return CookedFish
	case 351:
		return InkSack
	case 352:
		return Bone
	case 353:
		return Sugar
	case 354:
		return Cake
	case 355:
		return Bed
	case 356:
		return Diode
	case 357:
		return Cookie
	case 358:
		return Map
	case 359:
		return Shears
	case 2256:
		return GoldRecord
	case 2257:
		return GreenRecord
	}

	return Air
}
