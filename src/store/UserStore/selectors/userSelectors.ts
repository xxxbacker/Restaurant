import { UserStore } from '../store/UserStore.ts';

export const getUser = (state: UserStore) => state?.user
