import { Link } from "react-router";

function App() {
    return (
        <>
            <h1>Welcome to New Gopher's wet dream</h1>
            <p>
                This is a basic website created using golang and react, its
                purely raw. Created on 20/5/2025. I just want to learn golang,
                if you want please follow the
                <a
                    href="https://github.com/tausiq2003/newgopherswetdream"
                    target="_blank"
                >
                    Github
                </a>
                . We are making simple twitter like application, purely raw.
            </p>
            <br />
            <p>
                Get started <Link to={{ pathname: "/auth" }}>here</Link>.
            </p>
        </>
    );
}

export default App;
