type NamedAPIResource struct {
	name string
	url string
}

type Name struct {
	name string
	language Language
}

type Language struct {
	id int
	name string
	official bool
	iso639 string
	iso3166 string
	name []Name
}

type Generation struct {
	id int
	name string
	abilities
}

type VerboseEffect struct {
	effect string
	shortEffect string `json:"short_effect"`
	language Language
}

type AbilityEffectChange struct {
	effectEntries []Effect `json:"effect_entries"`
	versionGroup VersionGroup `json:"version_group"`
}

type AbilityFlavorText struct {
	flavorText string `json:"flavor_text"`
	language Language
	versionGroup VersionGroup `json:"version_group"`
}

type Ability struct {
	id int
	name string
	isMainSeries bool `json:"is_main_series"`
	generation Generation
	names []Name
	effectEntries []VerboseEffect `json:"effect_entries"`
	effectChanges []AbilityEffectChange `json"effect_changes"`
	flavorTextEntries AbilityFlavorText `json:"flavor_text_entries"`
	pokemon Pokemon
}

type AbilityPokemon struct {
	hidden bool `json:"is_hidden"`
	slot int
	ability Ability
}

type FormSprites struct {
	frontDefault string `json:"front_default"`
	frontShiny string `json:"front_shiny"`
	backDefault string `json:"back_default"`
	backShiny string `json:"back_shiny"`
}

type Description struct {
	description string
	language Language
}

type MoveLearnedMethod struct {
	id int
	name string
	descriptions []Description
	names []Name
	versionGroups []VersionGroup `json:"version_groups"`
}

type GrowthRateExperienceLevel struct {
	level int
	experience int
}

type GrowthRate struct {
	id int
	name string
	descriptions []Description
	levels []GrowthRateExperienceLevel
	pokemonSpecies PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpeciesDexEntyry struct {
	entryNumber int `json:"entry_number"`
	pokedex Pokedex
}

type EggGroup struct {
	id int
	name string
	names []Name
	pokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}

type PokemonColor struct {
	id int
	name string
	names []Name
	pokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}

type AwesomeName struct {
	awesomeName string `json:"awesome_name"`
	language Language
}

type PokemonShape struct {
	id int
	name string
	awesomeNames []AwesomeName `json:"awesome_names"`
	names []Name
	pokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}

type ItemFlingEffect struct {
	id int
	name string
	effectEntries []Effect
	items Item
}

type ItemAttribute struct {
	id int
	name string
	items Item
	names []Name
	descriptions []Description
}

type ItemPocket struct {
	id int
	name string
	categories []ItemCategory
	names []Name
}

type ItemCategory struct {
	id int
	name string
	items []Item
	names []Name
	pocket ItemPocket
}

type VersionGroupFlavorText struct {
	text string
	language Language
	versionGroup VersionGroup `json:"version_group"`
}

type GenerationGameIndex struct {
	gameIndex int `json:"game_index"`
	generation Generation
}

type ItemHolderPokemonVersionDetail struct {
	rarity int
	version Version
}

type ItemHolderPokemon struct {
	pokemon Pokemon
	versionDetails []ItemHolderPokemonVersionDetail
}

type ContestComboDetail struct {
	useBefore Move `json:"use_before"`
	useAfter Move `json:"use_after"`
}

type ContestComboSets struct {
	normal ContestComboDetail
	super ContestComboDetail
}

type BerryFlavor struct {
	id int
	name string
	names []Name
}

type ContestName struct {
	name string
	color string
	language Language
}

type ContestType struct {
	id int
	name string
	berryFlavor BerryFlavor `json:"berry_flavor"`
	names ContestName
}

type Move struct {
	id int
	name string
	accuracy int
	effectChance int `json:"effect_chance"`
	pp int
	priority int
	power int
	contestCombos ContestComboSets `json:"contest_combos"`
	contestType ContestType `json:"contest_type"`
}

type Machine struct {
	id int
	item Item
	move Move
	versionGroup VersionGroup `json:"version_group"`
}

type MachineVersionDetail struct {
	machine Machine
	versionGroup `json:"version_group"`
}

type ItemSprites struct {
}

type Item struct {
	id int
	name string
	cost in
	flingPower int `json:"fling_power"`
	flingEffect ItemFlingEffect `json:"fling_effect"`
	attributes []ItemAttribute
	category ItemCategory
	effectEntries VerboseEffect `json:"effect_entries"`
	flavorTextEntries VersionrGroupFlavorText `json:"flavor_text_entries"`
	gameIndices []GenerationGameIndex `json:"game_indices"`
	names []Name
	sprites []ItemSprites
	heldByPokemon ItemHolderPokemon `json:"held_by_pokemon"`
	babyTriggerFor EvolutionChain `json:"baby_trigger_for"`
	machines []MachineVersionDetail
}

type EvolutionChain struct {
	id int
	babyTriggerItem BabyTriggerItem `json:"baby_trigger_item"`
}

type PokemonHabitat struct {
}

type PalParkArea struct {
}

type PalParkEncounterArea struct {
	baseScore int `json:"base_score"`
	rate int
	area PalParPalParkkArea
}

type FlavorText struct {
}

type Genus struct {
	genus string
	language Language
}

type PokemonSpeciesVariety struct {
}

type PokemonSpecies struct {
	id int
	name string
	order int
	genderRate int `json:"gender_rate"`
	captureRate int `json:"capture_rate"`
	baseHappiness int `json:"base_happiness"`
	isBaby bool `json:"is_baby"`
	hatchCounter int `json:"hatch_counter"`
	hasGenderEffects bool `json:"has_gender_effects"`
	formsSwitchable bool `json:"forms_switchable"`
	growthRate GrowthRate `json:"growth_rate"`
}

type PokemonEntry struct {
	entryNumber int
	pokemonSpecies []PokemonSpecies `json:"pokemon_species"`
}

type Pokedex struct {
	id int
	name string
	isMainSeries bool `json:"is_main_series"`
	descriptions []Description
	names []Name
	pokemonEntries []PokemonEntry `json:"pokemon_entries"`
	region Region
	versionGroups []VersionGroup `json:"version_groups"`
}

type Region struct {
	id int
	name string
}

type Version struct {
	id int
	name string
	names []Name
	versionGroup VersionGroup `json:"version_group"`
}

type VersionGroup struct {
	id int
	name string
	order
	generation
	MoveLearnedMethods []MoveLearnedMethod `json:"move_learned_method"`
	pokedex []Pokedex
	regions []Region
	versions []Version
}

type PokemonForm struct {
	id int
	name string
	order int
	formOrder int `json:"form_order"`
	isDefault bool `json:"is_default"`
	isBattleOnly bool `json:"is_battle_only"`
	isMega bool `json:"is_mega"`
	formName string `json:"form_name"`
	pokemon Pokemon
	sprites []FormSprites
	versionGroup VersionGroup `json:"version_group"`
	names []Name
	formNames []Name `json:"form_names"`
}

type Pokemon struct {
	id int
	name string
	baseExperience int `json:"base_experience"`
	height int
	isDefault bool `json:"is_default"`
	order int
	weight int
	abilities []Ability
	forms []PokemonForm
}
