package main

type Sprites struct {
	BackDeault       string         `json:"back_default"`
	BackFemale       string         `json:"back_female"`
	BackShiny        string         `json:"back_shiny"`
	BackShinyFemale  string         `json:"back_shiny_female"`
	FrontDeault      string         `json:"front_default"`
	FrontFemale      string         `json:"front_female"`
	FrontShiny       string         `json:"front_shiny"`
	FrontShinyFemale string         `json:"front_shiny_female"`
	Other            Other          `json:"others"`
	Versions         VersionSprites `json:"versions"`
}

type VersionSprites struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationII   GenerationII   `json:"generation-ii"`
	GenerationIII  GenerationIII  `json:"generation-iii"`
	GenerationIV   GenerationIV   `json:"generation-iv"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVI   GenerationVI   `json:"generation-vi"`
	GenerationVII  GenerationVII  `json:"generation-vii"`
	GenerationVIII GenerationVIII `json:"generation-viii"`
}

type GenerationI struct {
	RedBlue map[string]string `json:"red-blue"`
	Yellow  map[string]string `json:"yellow"`
}

type GenerationII struct {
	Crystal map[string]string `json:"crystal"`
	Gold    map[string]string `json:"gold"`
	Silver  map[string]string `json:"silver"`
}

type GenerationIII struct {
	Emerald          map[string]string `json:"emerald"`
	FireredLeafgreen map[string]string `json:"firered-leafgreen"`
	RubySapphire     map[string]string `json:"ruby-sapphire"`
}

type GenerationIV struct {
	DiamonPearl         map[string]string `json:"diamond-pearl"`
	HeartGoldSoulSilver map[string]string `json:"heartgold-soulsilver"`
	Platinum            map[string]string `json:"platinum"`
}

type GenerationV struct {
	BlackWhite map[string]string `json:"black-white"`
}

type GenerationVI struct {
	OmegaRubyAlphaSapphire map[string]string `json:"omegaruby-alphasapphire"`
	XY                     map[string]string `json:"x-y"`
}

type GenerationVII struct {
	Icons             map[string]string `json:"icons"`
	UltraSunUltraMoon map[string]string `json:"ultra-sun-ultra-moon"`
}

type GenerationVIII struct {
	Icons map[string]string `json:"icons"`
}

type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
}

type DreamWorld struct {
	FrontDeault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
}

type Home struct {
	FrontDeault      string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type OfficialArtwork struct {
	FrontDeault string `json:"front_default"`
}

type Species struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Moves struct {
	Move                Move                  `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}

type VersionGroupDetails struct {
	LevelLearnedAt    int               `json:"level_learned_at"`
	MoveLearnedMethod MoveLearnedMethod `json:"move_learned_method"`
	VersionGroup      VersionGroup      `json:"version_group"`
}

type VersionGroup struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type MoveLearnedMethod struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Move struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GameIndices struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}

type Version struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Forms struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Ability struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}