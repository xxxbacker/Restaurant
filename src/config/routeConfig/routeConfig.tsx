import { RouteProps } from 'react-router-dom';
import { AuthPage } from 'pages';

export type AppRoutesProps = RouteProps & {
	authOnly?: boolean
}
export enum AppRoutes {
	MENU = 'menu',
	AUTH = 'auth',
	PROFILE = 'profile',
}

export const RoutePath: Record<AppRoutes, string> = {
	[AppRoutes.MENU]: '/menu',
	[AppRoutes.AUTH]: '/auth',
	[AppRoutes.PROFILE]: '/profile/' // + :id
}

export const routeConfig: Record<AppRoutes, AppRoutesProps> = {
	[AppRoutes.MENU]: {
		path: RoutePath.menu,
	},
	[AppRoutes.AUTH]: {
		path: RoutePath.auth,
		element: <AuthPage/>
	},
	[AppRoutes.PROFILE]: {
		path: RoutePath.profile,
	}
}
