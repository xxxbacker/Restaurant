import { classNames } from 'utils/classNames/classNames.ts';
import cls from './AuthPage.module.css'
import { Button, Form } from 'react-bootstrap';
import { UseAuthStore } from 'store/AuthStore/store/AuthStore.ts';
import { ChangeEvent, FormEvent, useState } from 'react';

interface AuthPageProps {
	className?: string;
}

const AuthPage = ({ className }: AuthPageProps) => {
	const [ isLogin, setIsLogin ] = useState(true)
	const { nickname, phone, password, setPassword, setPhone, setNickname, register, login } = UseAuthStore()
	const onNicknameChange = (event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
		setNickname(event.target.value)
	}
	const onPhoneChange = (event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
		setPhone(event.target.value)
	}
	const onPasswordChange = (event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
		setPassword(event.target.value)
	}
	const onSubmit = async (event: FormEvent<HTMLFormElement>) => {
		event.preventDefault()
		if (isLogin) {
			login(phone, password)
		} else {
			register(phone, password, nickname)
		}
	}
	return (
		<div className={classNames(cls.LoginPage, [ className, 'page' ], {})}>
			<div className={cls.container}>
				<Form onSubmit={onSubmit} className={cls.LoginForm}>
					<Form.Group controlId={'AuthForm.ControlPhone'}>
						<Form.Label className={cls.label}>Номер телефона</Form.Label>
						<Form.Control
							onChange={onPhoneChange}
							value={phone}
							type={'tel'}

							required></Form.Control>
					</Form.Group>
					<Form.Group controlId={'AuthForm.ControlPassword'}>
						<Form.Label className={cls.label}>Пароль</Form.Label>
						<Form.Control onChange={onPasswordChange} value={password} type={'password'} required></Form.Control>
					</Form.Group>
					{!isLogin && (
						<Form.Group controlId={'AuthForm.ControlNickname'}>
							<Form.Label className={cls.label}>Никнейм</Form.Label>
							<Form.Control
								onChange={onNicknameChange}
								value={nickname}
								required>
							</Form.Control>
						</Form.Group>
					)}
					<Button type="submit">{isLogin ? 'Войти' : 'Зарегистрироваться'}</Button>
				</Form>
				<div className={cls.swapWrapper}>
					<span>
						Нет аккаунта?
					</span>
					<span
					className={cls.swapLink}
					onClick={() => {
						setIsLogin(!isLogin)
						setPassword('')
					}}>{isLogin ? 'Зарегистрироваться' : 'Войти'}
					</span>
				</div>
			</div>
		</div>
	);
};
export default AuthPage
