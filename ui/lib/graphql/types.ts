export type Maybe<T> = T | null
export type InputMaybe<T> = Maybe<T>
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] }
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> }
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> }
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never }
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never }
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
	ID: { input: string; output: string }
	String: { input: string; output: string }
	Boolean: { input: boolean; output: boolean }
	Int: { input: number; output: number }
	Float: { input: number; output: number }
}

export type Comet = {
	__typename?: 'Comet'
	data: Scalars['String']['output']
	id: Scalars['ID']['output']
	patternId?: Maybe<Scalars['String']['output']>
}

export type Event = {
	__typename?: 'Event'
	page: Scalars['String']['output']
	userName: Scalars['String']['output']
}

export type Island = {
	__typename?: 'Island'
	comment: Scalars['String']['output']
	content: Scalars['String']['output']
	id: Scalars['ID']['output']
	planetId: Scalars['String']['output']
	title: Scalars['String']['output']
}

export type Mutation = {
	__typename?: 'Mutation'
	createComet?: Maybe<Scalars['String']['output']>
	createIsland?: Maybe<Scalars['String']['output']>
	createPattern?: Maybe<Scalars['String']['output']>
	createPlanet?: Maybe<Scalars['String']['output']>
	createStone?: Maybe<Scalars['String']['output']>
	deleteComet?: Maybe<Scalars['Boolean']['output']>
	deleteIsland?: Maybe<Scalars['Boolean']['output']>
	deletePlanet?: Maybe<Scalars['Boolean']['output']>
	deleteStone?: Maybe<Scalars['Boolean']['output']>
	linkComet?: Maybe<Scalars['Boolean']['output']>
	renamePlanet?: Maybe<Scalars['Boolean']['output']>
}

export type MutationCreateCometArgs = {
	data: Scalars['String']['input']
}

export type MutationCreateIslandArgs = {
	comment: Scalars['String']['input']
	content: Scalars['String']['input']
	planetId: Scalars['String']['input']
	title: Scalars['String']['input']
}

export type MutationCreatePatternArgs = {
	color?: InputMaybe<Scalars['String']['input']>
	title: Scalars['String']['input']
	traits: Array<TraitInput>
}

export type MutationCreatePlanetArgs = {
	comment: Scalars['String']['input']
	name: Scalars['String']['input']
}

export type MutationCreateStoneArgs = {
	data: Scalars['String']['input']
}

export type MutationDeleteCometArgs = {
	id: Scalars['ID']['input']
}

export type MutationDeleteIslandArgs = {
	id: Scalars['ID']['input']
}

export type MutationDeletePlanetArgs = {
	name: Scalars['String']['input']
}

export type MutationDeleteStoneArgs = {
	id: Scalars['ID']['input']
}

export type MutationLinkCometArgs = {
	id: Scalars['ID']['input']
	islandId: Scalars['String']['input']
}

export type MutationRenamePlanetArgs = {
	name: Scalars['String']['input']
	newName: Scalars['String']['input']
}

export type Pattern = {
	__typename?: 'Pattern'
	color?: Maybe<Scalars['String']['output']>
	id: Scalars['ID']['output']
	title: Scalars['String']['output']
	traits: Array<Trait>
}

export type Planet = {
	__typename?: 'Planet'
	comment: Scalars['String']['output']
	id: Scalars['ID']['output']
	name: Scalars['String']['output']
}

export type Query = {
	__typename?: 'Query'
	getComet?: Maybe<Comet>
	getCurrentSpace?: Maybe<Space>
	getIsland?: Maybe<Island>
	getPattern?: Maybe<Pattern>
	getPlanet?: Maybe<Planet>
	getStone?: Maybe<Stone>
	listComets: Array<Comet>
	listIslands: Array<Island>
	listPatterns: Array<Pattern>
	listPlanets: Array<Planet>
	listStones: Array<Stone>
}

export type QueryGetCometArgs = {
	id?: InputMaybe<Scalars['ID']['input']>
}

export type QueryGetIslandArgs = {
	id?: InputMaybe<Scalars['ID']['input']>
}

export type QueryGetPatternArgs = {
	id?: InputMaybe<Scalars['ID']['input']>
}

export type QueryGetPlanetArgs = {
	id?: InputMaybe<Scalars['ID']['input']>
	name?: InputMaybe<Scalars['String']['input']>
}

export type QueryGetStoneArgs = {
	id?: InputMaybe<Scalars['ID']['input']>
}

export type QueryListIslandsArgs = {
	planetId: Scalars['ID']['input']
}

export type QueryListStonesArgs = {
	islandId: Scalars['ID']['input']
}

export type Space = {
	__typename?: 'Space'
	name: Scalars['String']['output']
}

export type Stone = {
	__typename?: 'Stone'
	data: Scalars['String']['output']
	id: Scalars['ID']['output']
	islandId?: Maybe<Scalars['String']['output']>
	patternId?: Maybe<Scalars['String']['output']>
}

export type Subscription = {
	__typename?: 'Subscription'
	subscribeEvent?: Maybe<Event>
}

export type Trait = {
	__typename?: 'Trait'
	defaultValue: Scalars['String']['output']
	path: Scalars['String']['output']
	required: Scalars['Boolean']['output']
	type: Scalars['String']['output']
}

export type TraitInput = {
	defaultValue: Scalars['String']['input']
	path: Scalars['String']['input']
	required: Scalars['Boolean']['input']
	type: Scalars['String']['input']
}

export type _Service = {
	__typename?: '_Service'
	sdl: Scalars['String']['output']
}
