# Firegate Alpha

This is an Alpine-based Firewall OS built for Homelabbers.

WARNING: This is only a firewall, it does not have DHCP or other routing features. Thus, an L3 Managed switch with ip address 172.16.0.2/30 is required for a complete network setup.

## FEATURES:
- **Traffic Shaping**: Built-in Traffic Shaping setup with **CAKE** (Common Applications Kept Enhanced) algorithm having *wash* and other features.
- **IDS/IPS**: Detect (or) Prevent Intrusion with Suricata configued out-of-the-box.
- **Tor routing**: Certain VLAN traffic can be routed through Tor for anonymity on the internet.
- **Unbound with Ad-blocking**: Unbound DNS with a blocklist for ads, trackers, malware and nsfw.
- **Easy-to-use Web GUI**: Web GUI designed for easy configuration without the need to edit configuration files manually.

## ADVANTAGES:
- **Standalone Firewall**: As a separate L3 Switch will handle DHCP, Inter-VLAN Routing, etcetra, Firegate will be efficient at its core purpouse.
- **Homelabber-Centric**: Built with the homelabbers in mind.
- **Minimal Setup**: Core features are already setup for you.
- **Secure**: Definetly more secure than a standard home router.
- **Super Lightweight**: Alpine Linux is the base for this OS.

Core features developed. Web GUI still in development.
