package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/xml"
	"fmt"
	"log"
)

type JobConfig struct {
	XMLName xml.Name `xml:"job_config"`
	Name    string   `xml:"name"`
	Command string   `xml:"command"`
	Timeout int      `xml:"timeout_seconds"`
}

func main() {
	originalJob := JobConfig{
		Name: "Daily Backup",
		Command: "tar -czf backup.tar.gz /data",
		Timeout: 3600,
	}
	

	fmt.Println("=== 1. GOB (GO Binary) SERIALIZATION ===")
	var gobBuffer bytes.Buffer
	gobEncoder := gob.NewEncoder(&gobBuffer)
	if err := gobEncoder.Encode(originalJob); err != nil {
		log.Fatalf("GOB Encode Error: %v", err)
	}
	fmt.Printf("[GOB] Encoded binary length: %d bytes\n", gobBuffer.Len())
	fmt.Printf("[GOB] Raw data: %x\n", gobBuffer.Bytes())
	var decodedGobJob JobConfig
	gobDecoder := gob.NewDecoder(&gobBuffer)
	if err := gobDecoder.Decode(&decodedGobJob); err != nil {
		log.Fatalf("GOB Decode Error: %v", err)
	};fmt.Printf("[GOB] Decoded Struct: %+v\n\n", decodedGobJob)


	fmt.Println("=== 2. XML SERIALIZATION ===")
	xmlBytes, err := xml.MarshalIndent(originalJob, "", " ")
	if err != nil {
		log.Fatalf("XML Encode Error: %v", err)
	}
	xmlOutput := xml.Header + string(xmlBytes)
	fmt.Printf("[XML] Encoded Data:\n%s\n", xmlOutput)
	var decodedXMLJob JobConfig
	if err := xml.Unmarshal(xmlBytes, &decodedXMLJob); err != nil {
		log.Fatalf("XML Decode Error: %v", err)
	}; fmt.Printf("[XML] Decoded Struct: %+v\n\n", decodedXMLJob)

	
	fmt.Println("=== 3. BASE64 ENCODING ===")
	secretCommand := "sudo rm -rf /"
	encodedB64 := base64.StdEncoding.EncodeToString([]byte(secretCommand))
	fmt.Printf("[BASE64] Encoded Secret: %s\n", encodedB64)
	decodedB64, err := base64.StdEncoding.DecodeString(encodedB64)
	if err != nil {
		log.Fatalf("Base64 Decode Error: %v", err)
	}; fmt.Printf("[BASE64] Decoded secret: %s\n", string(decodedB64))
}