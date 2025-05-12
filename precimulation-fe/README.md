download node and npm

# Due to `npm install` failing, I ran these.
# https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_execution_policies?view=powershell-7.5
Get-ExecutionPolicy -List
Set-ExecutionPolicy -ExecutionPolicy Unrestricted -Scope CurrentUser

npm create vite@latest frontend --template react
  cd frontend
  npm install
  npm run dev
http://localhost:5173/
Create the dockerignore file is IMPORTANT!

Set-ExecutionPolicy -ExecutionPolicy Undefined -Scope CurrentUser

=====

# docker build --progress=plain --no-cache -t my-react-frontend .
docker build -t my-react-frontend .
docker run -p 80:80 my-react-frontend

=====

Set-ExecutionPolicy -ExecutionPolicy Unrestricted -Scope CurrentUser

npm install axios

Set-ExecutionPolicy -ExecutionPolicy Undefined -Scope CurrentUser
