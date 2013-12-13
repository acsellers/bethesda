package esm

type ESM struct {
	IsMasterFile bool
	MasterFile   string
	FileVersion  int
	Author       string
	Extra        []Attribute

	// Actor Data
	BodyPartDatas []BodyPartData `esm:"BPTD"`
	Challenges    []Challenge    `esm:"CHAL"`
	Classes       []Class        `esm:"CLAS"`
	EyesData      []Eyes         `esm:"EYES"`
	Factions      []Faction      `esm:"FACT"`
	HeadParts     []HeadPart     `esm:"HDPT"`
	AIPackages    []AIPackage    `esm:"PACK"`
	Perks         []Perk         `esm:"PERK"`
	Quests        []Quest        `esm:"QUST"`
	Races         []Race         `esm:"RACE"`
	Ragdolls      []Ragdoll      `esm:"RGDL"`
	Reputations   []Reputation   `esm:"REPU"`
	VoiceTypes    []VoiceType    `esm:"VTYP"`

	// Actors
	Creatures         []Creature         `esm:"CREA"`
	LeveledCreatures  []LeveledCreature  `esm:"LVLC"`
	LeveledCharactesr []LeveledCharacter `esm:"LVLN"`
	NPCs              []Npc              `esm:"NPC_"`
	TalkingActivators []TalkingActivator `esm:"TACT"`

	// Audio
	AcousticSpaces           []AcousticSpace           `esm:"ASPC"`
	MediaLocationControllers []MediaLocationController `esm:"ALOC"`
	MediaSets                []MediaSet                `esm:"MSET"`
	MusicTypes               []MusicType               `esm:"MUSC"`
	Sounds                   []Sound                   `esm:"SOUN"`

	// Game Effects
	Spells []Spell `esm:"SPEL"`
	// Derived Things
	// Abilities    []Ability
	// ActorEffects []ActorEffect
	// Addictions   []Addiction
	// Diseases     []Disease
	// Poisons      []Poison
	AmmoEffects  []AmmoEffect  `esm:"AMEF"`
	BaseEffects  []BaseEffects `esm:"MGEF"`
	Foods        []Food        `esm:"ALCH"`
	Enchantments []Enchantment `esm:"ENCH"`

	// Items
	Ammunitions   []Ammunition   `esm:"AMMO"`
	Armors        []Armor        `esm:"ARMO"`
	ArmorAddons   []ArmorAddon   `esm:"ARMA"`
	Books         []Book         `esm:"BOOK"`
	CaravanCards  []CaravanCard  `esm:"CCRD"`
	CaravanMoneys []CaravanMoney `esm:"CMNY"`
	ItemMods      []ItemMod      `esm:"IMOD"`
	Keys          []Key          `esm:"KEYM"`
	LeveledItems  []LeveledItem  `esm:"LVLI"`
	MiscItems     []MiscItem     `esm:"MISC"`
	Notes         []Note         `esm:"NOTE"`
	Weapons       []Weapon       `esm:"WEAP"`

	// Miscellaneous
	AnimationObjects []AnimationObject `esm:"ANIO"`
	CaravanDecks     []CaravanDeck     `esm:"CDCK"`
	Casinos          []Casino          `esm:"CSNO"`
	CasinoChips      []CasinoChip      `esm:"CHIP"`
	CombatStyles     []CombatStyle     `esm:"CSTY"`
	FormLists        []FormList        `esm:"FLST"`
	GlobalVars       []GlobalVar       `esm:"GLOB"`
	IdleMarkers      []IdleMarker      `esm:"IDLM"`
	LandTextures     []LandTexture     `esm:"LTEX"`
	LoadScreens      []LoadScreen      `esm:"LSCR"`
	LoadScreenTypes  []LoadScreenType  `esm:"LSCT"`
	MenuIcons        []MenuIcon        `esm:"MICN"`
	Messages         []Message         `esm:"MESG"`
	Recipes          []Recipe          `esm:"RCPE"`
	RecipeCategories []RecipeCategory  `esm:"RCCT"`
	Scripts          []Script          `esm:"SCPT"`
	TextureSets      []TextureSet      `esm:"TXST"`
	WaterTypes       []WaterType       `esm:"WATR"`

	// Special Effects
	AddOnNodes          []AddOnNode          `esm:"ADDN"`
	CameraShots         []CameraShot         `esm:"CAMS"`
	Debrises            []Debris             `esm:"DEBR"`
	EffectShaders       []EffectShader       `esm:"EFSH"`
	Explosions          []Explosion          `esm:"EXPL"`
	ImageSpaces         []ImageSpace         `esm:"IMGS"`
	ImageSpaceModifiers []ImageSpaceModifier `esm:"IMAD"`
	Impacts             []Impact             `esm:"IPCT"`
	ImpactDatas         []ImpactData         `esm:"IPDS"`
	Projectiles         []Projectile         `esm:"PROJ"`

	// World Objects
	Activators         []Activators       `esm:"ACTI"`
	Climates           []Climate          `esm:"CLMT"`
	Containers         []Container        `esm:"CONT"`
	Doors              []Door             `esm:"DOOR"`
	EncounterZones     []EncounterZone    `esm:"ECZN"`
	Furnitures         []Furniture        `esm:"FURN"`
	Grasses            []Grass            `esm:"GRAS"`
	Lights             []Light            `esm:"LIGH"`
	MovableStatics     []MovableStatic    `esm:"MSTT"`
	PlacableWaters     []PlacableWater    `esm:"PWAT"`
	Statics            []Static           `esm:"STAT"`
	StaticCollectionss []StaticCollection `esm:"SCOL"`
	Terminals          []Terminal         `esm:"TERM"`
	Trees              []Tree             `esm:"TREE"`
	Weathers           []Weather          `esm:"WTHR"`

	// Other World Records
	WorldSpaces       []WorldSpace       `esm:"WRLD" grouped:"true"`
	Regions           []Region           `esm:"REGN"`
	Cells             []Cell             `esm:"CELL" grouped:"true"`
	LightingTemplates []LightingTemplate `esm:"LGTM"`

	// Navmeshes
	NavigationInfos []NavigationInfo `esm:"NAVI"`

	// Other Character Data
	Hairs       []Hair       `esm:"HAIR"`
	ActorValues []ActorValue `esm:"AVIF"`
	Dialogues   []Dialogue   `esm:"DIAL" grouped:"true"`

	// Gameplay
	Radiations        []Radiation        `esm:"RADS"`
	Dehydrations      []Dehydration      `esm:"DEHY"`
	HungerStats       []HungerStat       `esm:"HUNG"`
	SleepDeprivations []SleepDeprivation `esm:"SLPD"`
	GameSettings      []GameSetting      `esm:"GMST"`
	IdleAnimations    []IdleAnimation    `esm:"IDLE"`
	CameraPaths       []CameraPath       `esm:"CPTH"`
	DefaultObjects    []DefaultObject    `esm:"DOBJ"`

	// Note: DoNotCreateNewIngrediantsWeArentUsingThemInFallout
	// Unused in Fallout 3/NV
	Ingrediants   []Ingrediant   `esm:"INGR"`
	CameraObjects []CameraObject `esm:"COBJ"`
}
