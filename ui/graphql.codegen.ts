import { type CodegenConfig } from '@graphql-codegen/cli'

export default {
	overwrite: true,
	schema: '../gql/*.graphql',
	generates: {
		'lib/graphql/types.ts': {
			plugins: ['typescript'],
		},
	},
	hooks: {
		afterAllFileWrite: ['pnpm format'],
	},
} satisfies CodegenConfig
