const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export const registerUser = async (token: string, name: string) => {
  const response = await fetch(`${API_URL}/user/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`,
    },
    body: JSON.stringify({ name }),
  });
  
  if (!response.ok) {
    throw new Error("Failed to register user");
  }
  
  return response.json();
};
