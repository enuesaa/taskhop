import { get, mutate } from '$lib/graphql/client'
import type { MutationRunCmdArgs } from './types'

const createQuery = `mutation ($input: RunCmdInput!) {
  runCmd(input: $input)
}`
export const useRunCmd = () =>
	mutate<MutationRunCmdArgs>(createQuery, {
		usekey: 'runCmd',
	})
