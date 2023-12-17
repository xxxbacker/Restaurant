import { Container, Navbar } from 'react-bootstrap';
import { RoutePath } from 'config/routeConfig/routeConfig.tsx';
import { NavLink } from 'react-router-dom';
import cls from './Navbar.module.css'
import { classNames } from 'utils/classNames/classNames.ts';

export const AppNavbar = () => {
	return (
		<Navbar className="">
			<Container className={classNames('justify-content-end', [ cls.container ])}>
				<NavLink
					className={(navData) => classNames(cls.link, [], { [cls.active]: navData.isActive })}
					to={RoutePath.menu}
				>
					Меню
				</NavLink>
				<NavLink
					className={(navData) => classNames(cls.link, [], { [cls.active]: navData.isActive })}
					to={RoutePath.auth}
				>
					Авторизация
				</NavLink>
			</Container>
		</Navbar>
	);
};

