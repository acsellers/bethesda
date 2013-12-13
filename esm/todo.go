package esm

type BodyPartData struct {
	RecordData
}
type Challenge struct {
	RecordData
}
type Class struct {
	RecordData
}
type Eyes struct {
	RecordData
}
type Faction struct {
	RecordData
}
type HeadPart struct {
	RecordData
}
type AIPackage struct {
	RecordData
}
type Perk struct {
	RecordData
}
type Quest struct {
	RecordData
}
type Race struct {
	RecordData
}
type Ragdoll struct {
	RecordData
}
type Reputation struct {
	RecordData
}
type VoiceType struct {
	RecordData
}

// Actors
type Creature struct {
	RecordData
}
type LeveledCreature struct {
	RecordData
}
type LeveledCharacter struct {
	RecordData
}
type Npc struct {
	RecordData
}
type TalkingActivator struct {
	RecordData
}

// Audio
type AcousticSpace struct {
	RecordData
}
type MediaLocationController struct {
	RecordData
}
type MediaSet struct {
	RecordData
}
type MusicType struct {
	RecordData
}
type Sound struct {
	RecordData
}

// Game Effects
type Spell struct {
	RecordData
}

// Derived Things
// Abilities    []Ability
// ActorEffects []ActorEffect
// Addictions   []Addiction
// Diseases     []Disease
// Poisons      []Poison
type AmmoEffect struct {
	RecordData
}
type BaseEffects struct {
	RecordData
}
type Food struct {
	RecordData
}
type Enchantment struct {
	RecordData
}

// Items
type Ammunition struct {
	RecordData
}
type Armor struct {
	RecordData
}
type ArmorAddon struct {
	RecordData
}
type Book struct {
	RecordData
}
type CaravanCard struct {
	RecordData
}
type CaravanMoney struct {
	RecordData
}
type ItemMod struct {
	RecordData
}
type Key struct {
	RecordData
}
type LeveledItem struct {
	RecordData
}
type MiscItem struct {
	RecordData
}
type Note struct {
	RecordData
}
type Weapon struct {
	RecordData
}

// Miscellaneous
type AnimationObject struct {
	RecordData
}
type CaravanDeck struct {
	RecordData
}
type Casino struct {
	RecordData
}
type CasinoChip struct {
	RecordData
}
type CombatStyle struct {
	RecordData
}
type FormList struct {
	RecordData
}
type GlobalVar struct {
	RecordData
}
type IdleMarker struct {
	RecordData
}
type LandTexture struct {
	RecordData
}
type LoadScreen struct {
	RecordData
}
type LoadScreenType struct {
	RecordData
}
type MenuIcon struct {
	RecordData
}
type Message struct {
	RecordData
}
type Recipe struct {
	RecordData
}
type RecipeCategory struct {
	RecordData
}
type Script struct {
	RecordData
}
type TextureSet struct {
	RecordData
}
type WaterType struct {
	RecordData
}

// Special Effects
type AddOnNode struct {
	RecordData
}
type CameraShot struct {
	RecordData
}
type Debris struct {
	RecordData
}
type EffectShader struct {
	RecordData
}
type Explosion struct {
	RecordData
}
type ImageSpace struct {
	RecordData
}
type ImageSpaceModifier struct {
	RecordData
}
type Impact struct {
	RecordData
}
type ImpactData struct {
	RecordData
}
type Projectile struct {
	RecordData
}

// World Objects
type Activators struct {
	RecordData
}
type Climate struct {
	RecordData
}
type Container struct {
	RecordData
}
type Door struct {
	RecordData
}
type EncounterZone struct {
	RecordData
}
type Furniture struct {
	RecordData
}
type Grass struct {
	RecordData
}
type Light struct {
	RecordData
}
type MovableStatic struct {
	RecordData
}
type PlacableWater struct {
	RecordData
}
type Static struct {
	RecordData
}
type StaticCollection struct {
	RecordData
}
type Terminal struct {
	RecordData
}
type Tree struct {
	RecordData
}
type Weather struct {
	RecordData
}

// Other World Records
type WorldSpace struct {
	RecordData
}
type Region struct {
	RecordData
}
type Cell struct {
	RecordData
}
type LightingTemplate struct {
	RecordData
}

// Navmeshes
type NavigationInfo struct {
	RecordData
}

// Other Character Data
type Hair struct {
	RecordData
}
type ActorValue struct {
	RecordData
}
type Dialogue struct {
	RecordData
}

// Gameplay
type Radiation struct {
	RecordData
}
type Dehydration struct {
	RecordData
}
type HungerStat struct {
	RecordData
}
type SleepDeprivation struct {
	RecordData
}
type GameSetting struct {
	RecordData
}
type IdleAnimation struct {
	RecordData
}
type CameraPath struct {
	RecordData
}
type DefaultObject struct {
	RecordData
}

// Unused in Fallout3/NV
type Ingrediant struct {
	RecordData
}

// Unused in Fallout3/NV
type CameraObject struct {
	RecordData
}

// Note: DoNotCreateNewIngrediantsWeArentUsingThemInFallout
// Ingrediants []Ingrediant `esm:"INGR"`
// Note: Empty group record
// CameraObjects []CameraObject `esm:"COBJ"`
