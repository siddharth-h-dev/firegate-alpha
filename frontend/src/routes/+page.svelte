<script lang="ts">
	import { goto } from '$app/navigation';
	
	let username = $state('');
	let password = $state('');
	let errorMessage = $state('');
	
	async function handleLogin() {
		errorMessage =  '';
		try {
			const API_BASE = `http://${window.location.hostname}:8080`;
			const response = await fetch(`${API_BASE}/api/auth/login`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify({
					username: username,
					password: password
				})
			});
			
			if (response.status === 401) {
				errorMessage = "Invalid Username or Password. Please try again.";
			} else if (!response.ok) {
				alert(`Backend Error ${response.status}`);
				throw new Error(`Backend Error ${response.status}`);
			}
			
			const data = await response.json(); 
			if (data.success) {
				goto('/ui/core/dashboard');
			}
		} catch (err: any) {
			if (err instanceof TypeError && err.message.includes('fetch')) {
				alert('CRITICAL: Could not connect to the Firegate backend. Wait for it to initialize or login into shell to fix the issue.');
				console.error('CRITICAL: Could not connect to Backend.')
				return;
			}
			console.error(err);
		}
	}
</script>

<div class="login-page">
	<div class="login-card">

		<img src="/assets/logo.png" alt="Firegate Alpha Logo" class="logo" />
				
		<!-- Header -->
		<div class="login-header">
			<h1>Firegate Login</h1>
			<p>Firegate Alpha, abiding by DTyF</p>
		</div>	
		
		<!-- Error Banner -->
		{#if errorMessage}
			<div class="error-banner">
				{errorMessage}
			</div>
		{/if}
		
		<!-- Login Form -->
		<form onsubmit={handleLogin}>
			<div class="form-group">
				<label for="user">Username</label>
				<input
					id="user"
					type="text"
					bind:value={username}
					placeholder="e.g. admin"
					required
				/>
			</div>
			
			<div class="form-group">
				<label for="pass">Password</label>
				<input
					id="pass"
					type="password"
					bind:value={password}
					placeholder="********"
					required
				/>
			</div>
			
			<button type="submit">Login</button>
		</form>
	</div>
</div>

<style>
	:global(html), :global(body) {
		--bg-page: #000000;
		--bg-card: #0a0a0a;
		--bg-input: #121212;
		--border-color: #262626;
		--text-main: #ffffff;
		--text-muted: #a3a3a3;
		--accent: #ff6600;
		--accent-hover: #cc5200;	
	}				
	
	.login-page {
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: var(--bg-page);
		color: var(--text-main);
		font-family: monospace;
		padding: 1rem;
	}
	
	.login-card {
		background-color: var(--bg-card);
		border: 1px solid var(--border-color);
		padding: 2rem;
		border-radius: 4px;
		width: 100%;
		max-width: 400px;
	}
	
	.login-header {
		text-align: center;
		margin-bottom: 1.5rem;
	}
	
	.login-header h1 {
		font-size: 1.3rem;
		font-weight: 700;
		color: var(--accent);
		margin: 0 0 0.5rem 0;
	}

	.login-header p {
		font-size: 0.75rem;
		color: var(--text-muted);
		margin: 0;
	}
	
	.error-banner {
		background-color: #2a0808;
		border: 1px solid #7f1d1d;
		color: #f87171;
		font-size: 0.8rem;
		padding: 0.5rem;
		margin-bottom: 1rem;
		text-align: center;
		border-radius: 2px;
	}
	
	form {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}
	
	.form-group {
		display: flex;
		flex-direction: column;
		gap: 0.375rem;
	}
	
	.form-group label {
		font-size: 0.75rem;
		color: var(--text-muted);
	}
	
	.form-group input {
		background-color: var(--bg-input);
		border: solid 1px var(--border-color);
		color: var(--text-main);
		border-radius: 2px;
		font-size: 0.875rem;
		padding: 0.625rem 0.75rem;
	}
	
	.form-group input:focus {
		outline: none;
		border-color: var(--accent);
	}
	
	button {
		background-color: var(--accent);
		color: #000000;
		font-weight: 700;
		font-size: 0.875rem;
		padding: 0.625rem 1rem;
		border: none;
		border-radius: 2px;
		cursor: pointer;
		margin-top: 0.5rem;	
	}
	
	button:hover {
		background-color: var(--accent-hover);
	}
	
	.logo {
		width: 64px;
		height: auto;
		margin: 0 auto 1rem;
		display: block;
	}
</style>
