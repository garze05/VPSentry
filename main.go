package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {

	/** Flags **/
	var adminUser *string = flag.String("admin-user", "ubuntu", "Specify the user's name for the admin account")
	
	flag.Parse()

	isAdmin, _ := UserIsAdmin(*adminUser)

	if isAdmin {
		fmt.Printf("El usuario %s es administrador\n", *adminUser)
	}

	
	// var archivo string = "test.txt"
	// resultado, err := existeArchivo(archivo)
	// fmt.Println(resultado, err)
	// err := prepararArchivo(*archivo)

	// if err != nil { fmt.Printf("Error: %v", err) }
}

func UserExists(name string) (bool, error) {
	// Ideally, this should be implemented using os/user package
	cmd := exec.Command("getent", "passwd", name)

	out, err := cmd.CombinedOutput()

if err != nil {
		var exitErr *exec.ExitError

		// try to extract exitErr safely
		if errors.As(err, &exitErr) {
			// 2 means user was not found using getent passwd <user>
			if exitErr.ExitCode() == 2 {
				return false, nil
			}
		}
		
		log.Printf("Error ejecutando el comando: %v\nSalida: %s", err, string(out))
		return false, err
	}

	// No error
	return true, nil
}

func UserIsAdmin(name string) (bool, error) {
	exists, err := UserExists(name)

	if err != nil {
		// log.Printf("Error ejecutando el comando: %v\nSalida: %s\n", err, string(out))
		return false, err
	}

	if !exists { return false, nil }

	cmd := exec.Command("id", "-nG", name)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return false, err
	}
	// Uid == "0" is root
	var searchGroup string = "sudo"

	if strings.Contains(string(out), searchGroup) {
		return true, nil
	}

	return false, nil
}

// func existeArchivo(ruta string) (bool, error) {
// 	fileInfo, err := os.Stat(ruta)

// 	if err == nil {
// 		if fileInfo.IsDir() {
// 			return false, fmt.Errorf("la ruta %s es un directorio", ruta)
// 		}

// 		return true, nil
// 	}

// 	if os.IsNotExist(err) {
// 		return false, nil
// 	}

// 	return false, err
// }

// func prepararArchivo(ruta string) error {
// 	existe, err := existeArchivo(ruta)

// 	if err != nil { return err }

// 	indice := strings.LastIndex(ruta, "/")
// 	archivo := ruta[indice+1:]

// 	if existe == false {
// 		file, err := os.OpenFile(ruta, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)

// 		if err != nil { return err }
		
// 		closeErr := file.Close()

// 		if closeErr != nil { return closeErr }

// 		fmt.Printf("Archivo \"%s\" creado con exito", archivo)
// 		return nil
// 	}

// 	fmt.Printf("El archivo \"%s\" ya existe. No se modificó", archivo)
// 	return nil
// }