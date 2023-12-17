import { create } from 'zustand';
import { User } from 'store/UserStore/store/UserStore.ts';
import { immer } from 'zustand/middleware/immer';
import { createSelectors } from 'utils/createSelectors/createSelectrors.ts';
import { devtools, persist } from 'zustand/middleware';
import { $host } from 'config/apiConfig/apiConfg.ts';

export interface AuthStore extends User {
	setNickname: (value: string) => void
	setPassword: (value: string) => void
	setPhone: (value: string) => void
	login: (phone: string, password: string) => void
	register: (phone: string, password: string, nickname: string) => void
}

const UseAuthStoreBase = create<AuthStore>()(devtools(immer(set => ({
	nickname: '',
	password: '',
	phone: '',
	setNickname: (value: string) => set((state: AuthStore) => {
		state.nickname = value
	}),
	setPassword: (value: string) => set((state) => {
		state.password = value
	}),
	setPhone: (value: string) => set((state) => {
		state.phone = value
	}),

	login: async (phone: string, password) => {
		const response = await $host.post('sign-in', {
			phone,
			password
		})
		console.log(response)
	},
	register: async (phone: string, password: string, nickname: string) => {
		const response = await $host.post('/sign-up', {
			phone, password, nickname
		})
		console.log(response)
	}
}))))
persist(UseAuthStoreBase, { name: 'UseAuthStoreBase' })
devtools(UseAuthStoreBase)
export const UseAuthStore = createSelectors(UseAuthStoreBase)
