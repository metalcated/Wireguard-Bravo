package main

import (
    "log"
    "os/exec"
    "strconv"

    "golang.zx2c4.com/wireguard/wgctrl"
)

type StatsStore struct {
    w *wgctrl.Client

    Alive bool
    Endpoint string
    Received string
    Sent string
}

func toReadable(bytes int64) string {
    const (
        KB = 1024 << (10 * iota) // 2^20
        MB
        GB
        TB
    )

    num := float64(bytes)

    switch {
    case bytes >= TB:
        return strconv.FormatFloat(num / TB, 'f', 2, 64) + " TiB"
    case bytes >= GB:
        return strconv.FormatFloat(num / GB, 'f', 2, 64) + " GiB"
    case bytes >= MB:
        return strconv.FormatFloat(num / MB, 'f', 2, 64) + " MiB"
    case bytes >= KB:
        return strconv.FormatFloat(num / KB, 'f', 2, 64) + " KiB"
    default:
        return strconv.FormatFloat(num, 'f', 2, 64) + " bytes"
    }
}

func (d *StatsStore) Update() *StatsStore {
    // init client
    if d.w == nil {
        c, err := wgctrl.New()
        if err != nil {
            log.Fatalf("Failed to open wireguardctrl: %v", err)
        }
        d.w = c
    }

    devs, err := d.w.Devices()
    if err != nil {
        log.Fatalf("Failed to get devices: %v", err)
    }

    if devs == nil {
        d.Alive = false
        return d
    }

    // TODO: to be rewritten:

    peers := devs[0].Peers // we only expect 1 wireguard connection
    if peers == nil {
        d.Alive = false
        return d
    }

    peer := peers[0] // and only one vpn server

    d.Alive    = true
    d.Endpoint = peer.Endpoint.IP.String()
    d.Received = toReadable(peer.ReceiveBytes)
    d.Sent     = toReadable(peer.TransmitBytes)

    return d
}

func (d *StatsStore) Close() {
    if err := d.w.Close(); err != nil {
        log.Println("Error closing wireguardctrl watcher:", err)
    }
}

// TODO: Buggy
func toggleConnection(disconnect bool) {
    var arg string

    if disconnect {
        arg = "down"
    } else {
        arg = "up"
    }

    if err := exec.Command("wg-quick", arg, conf.WgConfig).Run(); err != nil {
        log.Println("Error running wg-quick:", err)
    }
}
