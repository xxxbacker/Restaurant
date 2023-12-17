import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import { createSelectors } from 'utils/createSelectors/createSelectrors.ts';

export interface User {
	nickname: string,
	password: string,
	phone: string
}
export interface UserStore {
	user?: User
}
export const  UseUserStoreBase = create<UserStore>()(immer(set => ({
	setUser: (user: User) => set((state) => state.user = user),
})))

export const UseUserStore = createSelectors(UseUserStoreBase)
