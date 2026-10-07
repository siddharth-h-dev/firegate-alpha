<script>
	let openMenu = $state("");
	
	function toggleMenu(menuName) {
		openMenu = openMenu == menuName ? "" : menuName;
	}
</script>

<header class="header">
	<div class="header-left">
		<img src="/assets/logo.png" alt="Firegate Alpha Logo" class="logo" />
		<h1>Firegate Alpha</h1>
	</div>
	<div class="header-right">
		<button onclick={handleLogout} class="logout-btn">Logout</button>
	</div>
</header>

<main class="page">
	<aside class="sidebar">
		<nav class="sidebar-menu">
			<div class="tree-node" class:is-open={openMenu == 'lobby'}>
				<button class="tree-toggle" aria-expanded="false" onclick={() => toggleMenu('lobby')}>Lobby</button>
				<div class="tree-submenu">
					<div class="submenu-inner">
						<a href="/ui/core/dashboard" class="submenu-item">Dashboard</a>
						<a href="/ui/core/license" class="submenu-item">License</a>
					</div>
				</div>
			</div>
			<div class="tree-node" class:is-open={openMenu == 'system'}>
				<button class="tree-toggle" aria-expanded="false" onclick={() => toggleMenu('system')}>System</button>
				<div class="tree-submenu">
					<div class="submenu-inner">
						<a href="/ui/system/access" class="submenu-item">Access</a>
						<a href="/ui/system/firmware" class="submenu-item">Firmware</a>
						<a href="/ui/system/interfaces" class="submenu-item">Interfaces</a>
						<a href="/ui/system/settings" class="submenu-item">Settings</a>
					</div>
				</div>
			</div>
			<div class="tree-node" class:is-open={openMenu == 'gitops'}>
				<button class="tree-toggle" aria-expanded="false" onclick={() => toggleMenu('gitops')}>GitOps</button>
				<div class="tree-submenu">
					<div class="submenu-inner">
						<a href="/ui/git/repos/nftables" class="submenu-item">Nftables Repo</a>
						<a href="/ui/git/repos/suricata" class="submenu-item">Suricata Repo</a>
						<a href="/ui/git/repos/system" class="submenu-item">System Repo</a>
						<a href="/ui/git/repos/tor" class="submenu-item">Tor Repo</a>
						<a href="/ui/git/repos/unbound" class="submenu-item">Unbound Repo</a>
					</div>
				</div>
			</div>
			<div class="tree-node" class:is-open={openMenu == 'services'}>
				<button class="tree-toggle" aria-expanded="false" onclick={() => toggleMenu('services')}>Services</button>
				<div class="tree-submenu">
					<div class="submenu-inner">
						<a href="/ui/services/nftables" class="submenu-item">Nftables Firewall</a>
						<a href="/ui/services/suricata" class="submenu-item">Suricata IDS/IPS</a>
						<a href="/ui/services/unbound" class="submenu-item">Unbound DNS</a>
						<a href="/ui/services/tor" class="submenu-item">Tor</a>
					</div>
				</div>
			</div>
			<div class="tree-node" class:is-open={openMenu == 'power'}>
				<button class="tree-toggle" aria-expanded="false" onclick={() => toggleMenu('power')}>Power</button>
				<div class="tree-submenu">
					<div class="submenu-inner">
						<a href="/ui/system/reboot" class="submenu-item">Reboot</a>
						<a href="/ui/system/poweroff" class="submenu-item">Power off</a>
					</div>
				</div>
			</div>
		</nav>
	</aside>
	<section class="content">
	</section>
</main>
<style>
	:global(html), :global(body) {
		--bg-page: #000000;
		--bg-page2: #0a0a0a;
		--bg-page3: #121212;
		--border-color: #262626;
		--text-main: #ffffff;
		--text-muted: #a3a3a3;
		--accent: #ff6600;
		--accent-hover: #cc5200;
		--sidebar-text: #ffffff;
	}				
	
	.page {
		min-height: calc(100vh - 72px);
		min-width: 100vw;
		display: flex;
		gap: 30px;
		overflow: hidden;
		background: var(--bg-page);
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		background: var(--bg-page2);
		padding: 15px 30px;
	}
	
	.header-left {
		display: flex;
		align-items: center;
		gap: 15px;
	}
	
	.logo {
		width: 64px;
		height: auto;
		object-fit: contain;
	}
	
	h1 {
		margin: 0;
		font-size: 24px;
		color: var(--text-main);
	}
	
	.header-right {
		display: flex;
		align-items: center;
	}
	
	.logout-btn {
		background: var(--accent);
		color: #000000;
		font-weight: 700;
		font-size: 0.875rem;
		padding: 0.625rem 1rem;
		border: none;
		border-radius: 2px;
		cursor: pointer;
	}
	
	.logout-btn:hover {
		background: var(--accent-hover);
	}
	
	.sidebar {
		background: var(--bg-page2);
		width: 15rem;
		flex-shrink: 0;
	}
	.content {
		flex-grow: 1;
		padding: 2rem;
		overflow-y: auto;
	}

	.sidebar-menu {
		padding: 0;	
	}

	
	.tree-toggle {
		display: flex;
		justify-content: space-between;
		align-items: center;
		width: 100%;
		padding: 0.25rem 0.25rem;
		background: none;
		border: none;
		color: var(--sidebar-text);
		text-decoration: none;
		text-align: left;
		font-size: 1rem;
		cursor: pointer;
		box-sizing: border-box;
	}
	.tree-toggle:hover {
		background: var(--bg-page3);
	}
	
	.tree-submenu {
		display: grid;
		grid-template-rows: 0fr;
		transition: grid-template-rows 0.2s ease-out;
		overflow: hidden;
	}
	
	.tree-node.is-open .tree-submenu {
		grid-template-rows: 1fr;
	}
	
	.submenu-item {
		min-height: 0;
		display: block;
		padding: 0.25rem 1.5rem 0.25rem 2.5rem;
		color: var(--sidebar-text);
		text-decoration: none;
		transition: padding 0.2s ease-out;
	}
	
	.submenu-inner {
		overflow: hidden;
	}
	
	.submenu-item:hover {
		background-color: var(--bg-page3);
	}
</style>
