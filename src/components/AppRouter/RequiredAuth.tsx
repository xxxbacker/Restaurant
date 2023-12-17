import { Navigate, useLocation } from 'react-router-dom';
import { RoutePath } from 'config/routeConfig/routeConfig.tsx';
import { UseUserStore } from 'store/UserStore/store/UserStore.ts';

export function RequireAuth({ children }: { children: JSX.Element }) {
	const auth = UseUserStore.use.user?.();
	const location = useLocation();

	if (!auth) {
		return <Navigate to={RoutePath.menu} state={{ from: location }} replace />;
	}

	return children;
}
