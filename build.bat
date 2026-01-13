@echo off
REM build.bat - Build script for Word of God Gabriel Compiler (Windows)

setlocal enabledelayedexpansion

set BINARY_NAME=gabriel
set MAIN_FILE=main.go
set BUILD_DIR=build
set VERSION=1.0.0

if "%1%"=="" (
    call :build_current
    goto end
)

if /i "%1%"=="windows" (
    call :build_windows
    goto end
)

if /i "%1%"=="linux" (
    call :build_linux
    goto end
)

if /i "%1%"=="macos" (
    call :build_macos
    goto end
)

if /i "%1%"=="all" (
    call :build_all
    goto end
)

if /i "%1%"=="clean" (
    call :clean
    goto end
)

if /i "%1%"=="install" (
    call :install
    goto end
)

if /i "%1%"=="help" (
    call :print_help
    goto end
)

call :print_help
goto end

:print_help
echo Word of God - Gabriel Compiler Build Script
echo.
echo Usage: build.bat [OPTION]
echo.
echo Options:
echo   windows     Build for Windows (64-bit)
echo   linux       Build for Linux (64-bit)
echo   macos       Build for macOS (64-bit)
echo   all         Build for all platforms
echo   clean       Remove build artifacts
echo   install     Build and install to system
echo   help        Show this help message
echo.
echo Examples:
echo   build.bat              # Build for Windows
echo   build.bat all          # Build for all platforms
echo   build.bat install      # Build and install
goto :eof

:build_current
echo Building Gabriel for Windows...
if not exist "%BUILD_DIR%" mkdir "%BUILD_DIR%"
go build -v -o "%BUILD_DIR%\%BINARY_NAME%.exe" "%MAIN_FILE%"
if !errorlevel! equ 0 (
    echo Build complete: %BUILD_DIR%\%BINARY_NAME%.exe
) else (
    echo Build failed with error code !errorlevel!
    exit /b 1
)
goto :eof

:build_windows
echo Building Gabriel for Windows...
if not exist "%BUILD_DIR%" mkdir "%BUILD_DIR%"
setlocal
set GOOS=windows
set GOARCH=amd64
go build -v -o "%BUILD_DIR%\%BINARY_NAME%.exe" "%MAIN_FILE%"
endlocal
if !errorlevel! equ 0 (
    echo Build complete: %BUILD_DIR%\%BINARY_NAME%.exe
) else (
    echo Build failed with error code !errorlevel!
    exit /b 1
)
goto :eof

:build_linux
echo Building Gabriel for Linux...
if not exist "%BUILD_DIR%" mkdir "%BUILD_DIR%"
setlocal
set GOOS=linux
set GOARCH=amd64
go build -v -o "%BUILD_DIR%\%BINARY_NAME%-linux" "%MAIN_FILE%"
endlocal
if !errorlevel! equ 0 (
    echo Build complete: %BUILD_DIR%\%BINARY_NAME%-linux
) else (
    echo Build failed with error code !errorlevel!
    exit /b 1
)
goto :eof

:build_macos
echo Building Gabriel for macOS...
if not exist "%BUILD_DIR%" mkdir "%BUILD_DIR%"
setlocal
set GOOS=darwin
set GOARCH=amd64
go build -v -o "%BUILD_DIR%\%BINARY_NAME%-macos" "%MAIN_FILE%"
endlocal
if !errorlevel! equ 0 (
    echo Build complete: %BUILD_DIR%\%BINARY_NAME%-macos
) else (
    echo Build failed with error code !errorlevel!
    exit /b 1
)
goto :eof

:build_all
echo Building Gabriel for all platforms...
call :build_windows
call :build_linux
call :build_macos
if !errorlevel! equ 0 (
    echo All platform builds complete
) else (
    echo One or more builds failed
    exit /b 1
)
goto :eof

:clean
echo Cleaning build artifacts...
if exist "%BUILD_DIR%" (
    rmdir /s /q "%BUILD_DIR%"
    echo Build directory removed
)
go clean
echo Clean complete
goto :eof

:install
echo Building Gabriel for Windows...
if not exist "%BUILD_DIR%" mkdir "%BUILD_DIR%"
go build -v -o "%BUILD_DIR%\%BINARY_NAME%.exe" "%MAIN_FILE%"
if !errorlevel! neq 0 (
    echo Build failed
    exit /b 1
)

echo Installing Gabriel to system...
setlocal enabledelayedexpansion
set "LOCAL_BIN=%USERPROFILE%\AppData\Local\Programs\Gabriel"
if not exist "!LOCAL_BIN!" mkdir "!LOCAL_BIN!"
copy "%BUILD_DIR%\%BINARY_NAME%.exe" "!LOCAL_BIN!\%BINARY_NAME%.exe"
if !errorlevel! equ 0 (
    echo Gabriel installed to !LOCAL_BIN!\%BINARY_NAME%.exe
    echo.
    echo Add !LOCAL_BIN! to your PATH environment variable to use gabriel from anywhere
    echo.
    echo To add to PATH:
    echo   1. Press Win+X and select "System"
    echo   2. Click "Advanced system settings"
    echo   3. Click "Environment Variables"
    echo   4. Under "User variables", click "New"
    echo   5. Variable name: PATH
    echo   6. Variable value: !LOCAL_BIN!
) else (
    echo Installation failed
    exit /b 1
)
goto :eof

:end
echo.
echo Done!
endlocal
exit /b 0