import { useState } from "react";

export default function Auth() {
    const [input, setInput] = useState("");
    const [pass, setPass] = useState("");
    const [loinp, setLoInp] = useState("");
    const [lopass, setLoPass] = useState("");

    return (
        <>
            <h3>hello welcome, continue below: </h3>
            <div id="auth">
                <form id="signup">
                    <p>Sign up:</p>
                    <div
                        id="inputandbutton"
                        style={{
                            display: "flex",
                            flexDirection: "column",
                            gap: "6px",
                            maxWidth: "min-content",
                        }}
                    >
                        <label htmlFor="username">
                            Username:
                            <input
                                name="uname"
                                type="text"
                                value={input}
                                onChange={(e) => {
                                    setInput(e.target.value);
                                }}
                            />
                        </label>
                        <label htmlFor="username">
                            Password:
                            <input
                                name="pass"
                                type="password"
                                value={pass}
                                onChange={(e) => {
                                    setPass(e.target.value);
                                }}
                            />
                        </label>
                        <button>Submit</button>
                    </div>
                </form>
                <form id="login">
                    <p>Login:</p>
                    <div
                        id="inputandbutton"
                        style={{
                            display: "flex",
                            flexDirection: "column",
                            gap: "6px",
                            maxWidth: "min-content",
                        }}
                    >
                        <label htmlFor="username">
                            Username:
                            <input
                                name="uname"
                                type="text"
                                value={loinp}
                                onChange={(e) => {
                                    setLoInp(e.target.value);
                                }}
                            />
                        </label>
                        <label htmlFor="username">
                            Password:
                            <input
                                name="pass"
                                type="password"
                                value={lopass}
                                onChange={(e) => {
                                    setLoPass(e.target.value);
                                }}
                            />
                        </label>
                        <button>Submit</button>
                    </div>
                </form>
            </div>
        </>
    );
}
