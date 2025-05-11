download node and npm

Due to `npm install` failing, I ran these.
https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_execution_policies?view=powershell-7.5
Get-ExecutionPolicy -List
Set-ExecutionPolicy -ExecutionPolicy Unrestricted -Scope CurrentUser

npm create vite@latest frontend --template react
  cd frontend
  npm install
  npm run dev

Set-ExecutionPolicy -ExecutionPolicy Undefined -Scope CurrentUser