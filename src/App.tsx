import { Suspense } from 'react'
import './App.css'
import { AppRouter } from 'components/AppRouter/AppRouter';
import { AppNavbar } from 'components/Navbar/Navbar.tsx';

function App() {
  return (
   <div className={'app'}>
     <AppNavbar/>
     <Suspense fallback={''}><AppRouter /></Suspense>
   </div>
  )
}

export default App
