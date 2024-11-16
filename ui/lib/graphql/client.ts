import { Client, createRequest, fetchExchange, type OperationResult, gql } from '@urql/svelte'
import { PUBLIC_GRAPHQL_ENDPOINT } from '$env/static/public'
import { createMutation, createQuery, useQueryClient } from '@tanstack/svelte-query'

export const client = new Client({
	url: PUBLIC_GRAPHQL_ENDPOINT,
	exchanges: [fetchExchange],
})

const calcCacheKey = (query: string, variables: Record<string, any>): string[] => {
	return [`${query}-${JSON.stringify(variables)}`]
}

type GetOptions = {
	vars: Record<string, any>
	usekey: string
	initialData: any
}
export const get = <T>(query: string, options: Partial<GetOptions> = {}) => {
	const vars = options.vars ?? {}
	const usekey = options.usekey ?? ''
	const initialData = options.initialData ?? undefined

	const creation = createQuery<T>({
		queryKey: calcCacheKey(query, vars),
		queryFn: async () => {
			const request = createRequest(gql(query), vars)
			const res = await client.executeQuery(request, {})
			if (res.data === undefined || res.data === null) {
				return res.data
			}
			if (res.data.hasOwnProperty(usekey)) {
				return res.data[usekey]
			}
			return {}
		},
		initialData: initialData,
	})

	return creation
}

type Vars = Record<string, any>
type MutateOptions = {
	usekey: string
}
export const mutate = <T extends Vars>(query: string, options: Partial<MutateOptions> = {}) => {
	const usekey = options.usekey ?? ''
	const queryClient = useQueryClient()

	const creation = createMutation({
		mutationFn: async (data: T) => {
			const request = createRequest(gql(query), data)
			const res = await client.executeMutation(request, {})
			if (res.data === undefined || res.data === null) {
				return res.data
			}
			await queryClient.invalidateQueries()
			if (res.data.hasOwnProperty(usekey)) {
				return res.data[usekey]
			}
			return {}
		},
	})

	return creation
}
