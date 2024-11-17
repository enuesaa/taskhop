import { get } from '$lib/graphql/client'
import type { Health } from './types'

const getQuery = `query {
	getHealth {
    ok
    code
    
	}
}`

export const getHealth = () =>
	get<Health>(getQuery, {
		usekey: 'getHealth',
	})
