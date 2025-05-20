import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { BrowserRouter, Route, Routes } from "react-router";
import Auth from "./components/auth.tsx";

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <BrowserRouter>
            <Routes>
                <Route path="/" element=<App /> />
                <Route path="/auth" element=<Auth /> />
            </Routes>
        </BrowserRouter>
    </StrictMode>,
);
