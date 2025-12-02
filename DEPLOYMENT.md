# Deployment Guide - Fly.io

This guide will help you deploy the "12 Days of Beckie" application to Fly.io.

## Prerequisites

1. A Fly.io account (sign up at https://fly.io/app/sign-up)
2. Credit card (required by Fly.io, ~$2-3/month cost)
3. flyctl CLI installed

## Step 1: Install flyctl CLI

### macOS (using Homebrew)
```bash
brew install flyctl
```

### macOS/Linux (using install script)
```bash
curl -L https://fly.io/install.sh | sh
```

### Windows (using PowerShell)
```powershell
pwsh -Command "iwr https://fly.io/install.ps1 -useb | iex"
```

## Step 2: Login to Fly.io

```bash
flyctl auth login
```

This will open a browser window for you to authenticate.

## Step 3: Launch the Application

From the project root directory, run:

```bash
flyctl launch
```

You'll be prompted with several questions:

1. **App Name**: Choose a unique name (e.g., `beckies-12-days`, `aob-2024`)
   - This will be your URL: `https://your-app-name.fly.dev`

2. **Region**: Choose closest to your girlfriend's location
   - `lhr` = London
   - `ams` = Amsterdam
   - `fra` = Frankfurt
   - `ewr` = New York
   - Full list: https://fly.io/docs/reference/regions/

3. **Postgres database**: Say **NO** (we don't need a database)

4. **Deploy now**: Say **YES**

The CLI will:
- Create the app on Fly.io
- Build your Docker image
- Deploy the application
- Give you a URL like `https://your-app-name.fly.dev`

## Step 4: Verify Deployment

Once deployed, visit your URL to verify it's working:
```bash
flyctl open
```

Or manually visit: `https://your-app-name.fly.dev`

## Step 5: (Optional) Add Custom Domain

If you have a custom domain:

1. Add the domain to your Fly.io app:
```bash
flyctl certs add yourdomain.com
flyctl certs add www.yourdomain.com
```

2. Add DNS records (from your domain registrar):
   - **A record**: `@` → IP address shown by Fly.io
   - **AAAA record**: `@` → IPv6 address shown by Fly.io
   - **CNAME record**: `www` → `your-app-name.fly.dev`

3. Wait for DNS propagation (5-30 minutes)

4. Verify SSL certificate:
```bash
flyctl certs show yourdomain.com
```

## Useful Commands

### View logs
```bash
flyctl logs
```

### Check app status
```bash
flyctl status
```

### SSH into the running app
```bash
flyctl ssh console
```

### Redeploy after making changes
```bash
flyctl deploy
```

### Scale resources (if needed)
```bash
flyctl scale memory 512  # Increase to 512MB
```

### Stop the app (to save costs)
```bash
flyctl scale count 0
```

### Restart the app
```bash
flyctl scale count 1
```

## Important Notes

### Timezone Configuration
The app is configured to use `Europe/London` timezone in `fly.toml`. This ensures the calendar days unlock at midnight UK time. If you need a different timezone, edit the `TZ` environment variable in `fly.toml`:

```toml
[env]
  TZ = "America/New_York"  # Change to your timezone
```

See: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

### Cost Optimization
The `fly.toml` is configured to automatically stop the VM when not in use and start it when accessed:
```toml
auto_stop_machines = "stop"
auto_start_machines = true
min_machines_running = 0
```

This keeps costs low (~$2-3/month) but means the first request after inactivity may take a few seconds.

### Updating Content
If you want to update the content files (day1.txt, day2.txt, etc.):

1. Edit the files locally
2. Commit to git
3. Run: `flyctl deploy`

## Troubleshooting

### "App name already taken"
Choose a different name when running `flyctl launch`.

### "Build failed"
Check the build logs. Common issues:
- Ensure all files are committed to git
- Check Dockerfile syntax
- Ensure `go.mod` is present

### "App not accessible"
- Check app status: `flyctl status`
- View logs: `flyctl logs`
- Ensure HTTP service is running on port 8080

### SSL/HTTPS issues with custom domain
- Wait 15-30 minutes for certificate provisioning
- Check certificate status: `flyctl certs show yourdomain.com`
- Verify DNS records are correct

## Support

- Fly.io Documentation: https://fly.io/docs/
- Fly.io Community Forum: https://community.fly.io/
- Project Issues: https://github.com/James-Francis-MT/aob/issues
