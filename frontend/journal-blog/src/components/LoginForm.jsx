import '../App.css'

function LoginForm() {
    return (
        <div>
            <form>
                <label for="username">Username</label>
                <input type="text" id="username" name="username" />
                <label for="password">Password</label>
                <input type="password" id="password" name="password" />
            </form>
            <button>Login</button>
        </div>
  );
}

export default LoginForm;