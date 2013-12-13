package esm

// All Records have certain fields in common, they are
// FormID (4byte identifier)
// RecordName (4byte string for type)
// RecordType (Go-specific constant for RecordName)
// EditorID (string name for the record)
type RecordData struct {
	RecordType
	FormID     int32
	Flags      [3]int32
	RecordName string
	EditorID   string
	Attrs      []Attribute
}

type Attribute struct {
	Name  string
	Value []byte
}

type RecordType int32

const (
	UNKNOWN RecordType = iota
	BODY_PART_DATA
	CHALLENGE
	CLASS
	EYES
	FACTION
	HEAD_PART
	AI_PACKAGE
	PERK
	QUEST
	RACE
	RAGDOLL
	REPUTATION
	VOICE_TYPE

	// Actors
	CREATURE
	LEVELED_CREATURE
	LEVELED_CHARACTER
	NPC
	TALKING_ACTIVATOR

	// Audio
	ACOUSTIC_SPACE
	MEDIA_LOCATION_CONTROLLER
	MEDIA_SET
	MUSIC_TYPE
	SOUND

	// Game Effects
	SPELL
	// Derived Things
	// Abilities    []Ability
	// ActorEffects []ActorEffect
	// Addictions   []Addiction
	// Diseases     []Disease
	// Poisons      []Poison
	AMMO_EFFECT
	BASE_EFFECTS
	FOOD
	ENCHANTMENT

	// Items
	AMMUNITION
	ARMOR
	ARMOR_ADDON
	BOOK
	CARAVAN_CARD
	CARAVAN_MONEY
	ITEM_MOD
	KEY
	LEVELED_ITEM
	MISC_ITEM
	NOTE
	WEAPON

	// Miscellaneous
	ANIMATION_OBJECT
	CARAVAN_DECK
	CASINO
	CASINO_CHIP
	COMBAT_STYLE
	FORM_LIST
	GLOBAL_VAR
	IDLE_MARKER
	LAND_TEXTURE
	LOAD_SCREEN
	LOAD_SCREEN_TYPE
	MENU_ICON
	MESSAGE
	RECIPE
	RECIPE_CATEGORY
	SCRIPT
	TEXTURE_SET
	WATER_TYPE

	// Special Effects
	ADDON_NODE
	CAMERA_SHOT
	DEBRIS
	EFFECT_SHADER
	EXPLOSION
	IMAGE_SPACE
	IMAGE_SPACE_MODIFIER
	IMPACT
	IMPACT_DATA
	PROJECTILE

	// World Objects
	ACTIVATORS
	CLIMATE
	CONTAINER
	DOOR
	ENCOUNTER_ZONE
	FURNITURE
	GRASS
	LIGHT
	MOVABLE_STATIC
	PLACABLE_WATER
	STATIC
	STATIC_COLLECTION
	TERMINAL
	TREE
	WEATHER

	// Other World Records
	WORLD_SPACE
	REGION
	CELL
	LIGHTING_TEMPLATE

	// Navmeshes
	NAVIGATION_INFO

	// Other Character Data
	HAIR
	ACTOR_VALUE
	DIALOGUE

	// Gameplay
	RADIATION
	DEHYDRATION
	HUNGER_STAT
	SLEEP_DEPRIVATION
	GAME_SETTING
	IDLE_ANIMATION
	CAMERA_PATH
	DEFAULT_OBJECT

	// Unused in Fallout3/NV
	INGREDIANT
	CAMERA_OBJECT
)
