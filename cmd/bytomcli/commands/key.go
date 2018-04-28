package commands

import (
	"encoding/hex"
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"

	"github.com/bytom/crypto/ed25519/chainkd"
	"github.com/bytom/util"
)

var createKeyCmd = &cobra.Command{
	Use:   "create-key <alias> <password>",
	Short: "Create a key",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var key = struct {
			Alias    string `json:"alias"`
			Password string `json:"password"`
		}{Alias: args[0], Password: args[1]}

		data, exitCode := util.ClientCall("/create-key", &key)
		if exitCode != util.Success {
			os.Exit(exitCode)
		}

		printJSON(data)
	},
}

var deleteKeyCmd = &cobra.Command{
	Use:   "delete-key <xpub> <password>",
	Short: "Delete a key",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		xpub := new(chainkd.XPub)
		if err := xpub.UnmarshalText([]byte(args[0])); err != nil {
			jww.ERROR.Println("delete-key:", err)
			os.Exit(util.ErrLocalExe)
		}

		var key = struct {
			Password string
			XPub     chainkd.XPub `json:"xpub"`
		}{XPub: *xpub, Password: args[1]}

		if _, exitCode := util.ClientCall("/delete-key", &key); exitCode != util.Success {
			os.Exit(exitCode)
		}
		jww.FEEDBACK.Println("Successfully delete key")
	},
}

var listKeysCmd = &cobra.Command{
	Use:   "list-keys",
	Short: "List the existing keys",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		data, exitCode := util.ClientCall("/list-keys")
		if exitCode != util.Success {
			os.Exit(exitCode)
		}

		printJSONList(data)
	},
}

var resetKeyPwdCmd = &cobra.Command{
	Use:   "reset-key-password <xpub> <old-password> <new-password>",
	Short: "Delete a key",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		xpub := new(chainkd.XPub)
		if err := xpub.UnmarshalText([]byte(args[0])); err != nil {
			jww.ERROR.Println("reset-key-password args not valid:", err)
			os.Exit(util.ErrLocalExe)
		}

		ins := struct {
			XPub        chainkd.XPub `json:"xpub"`
			OldPassword string       `json:"old_password"`
			NewPassword string       `json:"new_password"`
		}{XPub: *xpub, OldPassword: args[1], NewPassword: args[2]}

		if _, exitCode := util.ClientCall("/reset-key-password", &ins); exitCode != util.Success {
			os.Exit(exitCode)
		}
		jww.FEEDBACK.Println("Successfully reset key password")
	},
}

var signMsgCmd = &cobra.Command{
	Use:   "sign-message <xpub> <message> <password>",
	Short: "sign message to generate signature",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		xpub := chainkd.XPub{}
		if err := xpub.UnmarshalText([]byte(args[0])); err != nil {
			jww.ERROR.Println(err)
			os.Exit(util.ErrLocalExe)
		}

		var req = struct {
			RootXPub chainkd.XPub `json:"root_xpub"`
			Message  []byte       `json:"message"`
			Password string       `json:"password"`
		}{RootXPub: xpub, Message: []byte(args[1]), Password: args[2]}

		data, exitCode := util.ClientCall("/sign-message", &req)
		if exitCode != util.Success {
			os.Exit(exitCode)
		}
		printJSON(data)
	},
}

var verifyMsgCmd = &cobra.Command{
	Use:   "verify-message <xpub> <message> <signature>",
	Short: "verify signature for specified message",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		xpub := chainkd.XPub{}
		if err := xpub.UnmarshalText([]byte(args[0])); err != nil {
			jww.ERROR.Println(err)
			os.Exit(util.ErrLocalExe)
		}

		signature, err := hex.DecodeString(args[2])
		if err != nil {
			jww.ERROR.Println("convert signature error:", err)
			os.Exit(util.ErrLocalExe)
		}

		var req = struct {
			RootXPub  chainkd.XPub `json:"root_xpub"`
			Message   []byte       `json:"message"`
			Signature []byte       `json:"signature"`
		}{RootXPub: xpub, Message: []byte(args[1]), Signature: signature}

		data, exitCode := util.ClientCall("/verify-message", &req)
		if exitCode != util.Success {
			os.Exit(exitCode)
		}
		printJSON(data)
	},
}
