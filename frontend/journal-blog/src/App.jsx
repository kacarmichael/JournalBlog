import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import NavBar from './components/NavBar'
import LoginForm from './components/LoginForm'

function App() {

  return (
    <div>
      <header>
        <NavBar></NavBar>
        <LoginForm></LoginForm>
      </header>
    </div>
  )
}

export default App
