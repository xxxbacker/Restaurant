import { Route, Routes } from 'react-router-dom'
import { ReactElement, Suspense, useCallback } from 'react'
import { routeConfig } from 'config/routeConfig/routeConfig.tsx'
import { AppRoutesProps } from 'config/routeConfig/routeConfig'
import { RequireAuth } from './RequiredAuth'

export const AppRouter = (): ReactElement => {
    const renderWithWrapper = useCallback((route: AppRoutesProps) => {
        const element = <Suspense fallback={''}>{route.element}</Suspense>
        return (
          <Route
            key={route.path}
            path={route.path}
            element={route.authOnly ? <RequireAuth>{element}</RequireAuth> : element}
          />
        )
    }, [])

    return <Routes>{Object.values(routeConfig).map(renderWithWrapper)}</Routes>
}
