FROM hugomods/hugo:exts

COPY . /src

# Uncomment below to install dependencies via NPM, it depends on your site.
# RUN npm i

# Build site.
RUN hugo --minify 

