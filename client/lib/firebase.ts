import { initializeApp } from "firebase/app";
import { getAuth, GoogleAuthProvider } from "firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyDd1XJbKA_U-fwwVvkJR3op6k1cuRXbUAE",
  authDomain: "whisper-71362.firebaseapp.com",
  projectId: "whisper-71362",
  storageBucket: "whisper-71362.firebasestorage.app",
  messagingSenderId: "1066128573481",
  appId: "1:1066128573481:web:fc35b0a96618dbbce38838",
  measurementId: "G-CWR74DXJDN"
};

const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
export const googleProvider = new GoogleAuthProvider();
