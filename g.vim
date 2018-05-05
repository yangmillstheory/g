function! s:Google(type)
  let saved = @@
  if a:type ==# 'V'
    normal! `<V`>y
  elseif a:type ==# 'v'
    normal! `<v`>y
  elseif a:type ==# 'char'
    normal! `[v`]y
  else
    return
  endif
  " remove trailing newline
  "
  "   https://stackoverflow.com/a/6228454
  let @@ = substitute(strtrans(@@), '\^@', '', 'g')
  call system('g ' . shellescape(@@))
  let @@ = saved
endfunction

vnoremap <F2> :<c-u>call <SID>Google(visualmode())<cr>
nnoremap <F2> :set operatorfunc=<SID>Google<cr>g@
nnoremap <F2><F2> :silent !g <c-r><c-w><cr>
