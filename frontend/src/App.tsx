import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import Navbar from "./components/navbar/navbar";
import RegistrationForm from "./components/registration/registation";
import LoginForm, {useAuth} from "./components/login/login";
import Common from "./components/common/common";
import ProtectedNavbar from "./components/navbar/protected_navbar";
import MyFiles from "./components/my/my";
import SharedFiles from "./components/shared/shared";
import {ShareFiles} from "./components/share/share";
import {UploadFile} from "./components/upload/upload";
import {DeleteFiles} from "./components/delete/delete";

function App() {

    const [logged] = useAuth();

    return (
        <Router>
            {!logged && <Navbar/>}
            {logged && <ProtectedNavbar/>}
            <Routes>
                <Route path="/common" element={<Common/>} />
                <Route path="/login" element={<LoginForm/>}/>
                <Route path="/registration" element={<RegistrationForm/>}/>
                <Route path="/my" element={<MyFiles/>}/>
                <Route path="/shared" element={<SharedFiles/>}/>
                <Route path="/share" element={<ShareFiles/>}/>
                <Route path="/upload" element={<UploadFile/>}/>
                <Route path="/delete" element={<DeleteFiles/>}/>
            </Routes>
        </Router>

    );
}


export default App;