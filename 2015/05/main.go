package main

import (
	_ "embed"
	"github.com/harryalaw/advent-of-go/util"
	"regexp"
	"strings"
)

func solve(part func([]string) int) int {
	result := part(parseInput())
	println(result)
	return result
}

//go:embed data.txt
var input string

func parseInput() []string {
	return strings.Split(input, "\n")
}

func part1(input []string) int {
	niceCount := 0

	for _, line := range input {
		if isNiceString(line, false) {
			niceCount++
		}
	}

	return niceCount
}

func part2(input []string) int {
	niceCount := 0

	for _, line := range input {
		if isNiceString(line, true) {
			niceCount++
		}
	}
	return niceCount
}

func isNiceString(line string, isPart2 bool) bool {
	if isPart2 {
		return twoPairs(line) &&
			oneGap(line)
	}

	return atLeastThreeVowels(line) &&
		repeatedConsecutiveLetter(line) &&
		doesNotContainBadPairs(line)
}

func applyRegexMatch(line, pattern string, predicate func([]string) bool) bool {
	compiles := regexp.MustCompile(pattern)
	matches := compiles.FindAllString(line, -1)
	return predicate(matches)
}

func atLeastThreeVowels(line string) bool {
	vowels := regexp.MustCompile("a|e|i|o|u")
	matches := vowels.FindAllString(line, -1)
	return len(matches) >= 3
}

func repeatedConsecutiveLetter(line string) bool {
	// repeated := regexp.MustCompile("aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz")
	// matches := repeated.FindAllString(line, -1)
	// return len(matches) >= 1
	for index := 1; index < len(line); index++ {
		if line[index] == line[index-1] {
			return true
		}
	}
	return false
}

func doesNotContainBadPairs(line string) bool {
	badPairs := regexp.MustCompile("ab|cd|pq|xy")
	matches := badPairs.FindAllString(line, -1)
	return len(matches) == 0
}

func twoPairs(line string) bool {
	// twoPairs := regexp.MustCompile("aa.*aa|ab.*ab|ac.*ac|ad.*ad|ae.*ae|af.*af|ag.*ag|ah.*ah|ai.*ai|aj.*aj|ak.*ak|al.*al|am.*am|an.*an|ao.*ao|ap.*ap|aq.*aq|ar.*ar|as.*as|at.*at|au.*au|av.*av|aw.*aw|ax.*ax|ay.*ay|az.*az|ba.*ba|bb.*bb|bc.*bc|bd.*bd|be.*be|bf.*bf|bg.*bg|bh.*bh|bi.*bi|bj.*bj|bk.*bk|bl.*bl|bm.*bm|bn.*bn|bo.*bo|bp.*bp|bq.*bq|br.*br|bs.*bs|bt.*bt|bu.*bu|bv.*bv|bw.*bw|bx.*bx|by.*by|bz.*bz|ca.*ca|cb.*cb|cc.*cc|cd.*cd|ce.*ce|cf.*cf|cg.*cg|ch.*ch|ci.*ci|cj.*cj|ck.*ck|cl.*cl|cm.*cm|cn.*cn|co.*co|cp.*cp|cq.*cq|cr.*cr|cs.*cs|ct.*ct|cu.*cu|cv.*cv|cw.*cw|cx.*cx|cy.*cy|cz.*cz|da.*da|db.*db|dc.*dc|dd.*dd|de.*de|df.*df|dg.*dg|dh.*dh|di.*di|dj.*dj|dk.*dk|dl.*dl|dm.*dm|dn.*dn|do.*do|dp.*dp|dq.*dq|dr.*dr|ds.*ds|dt.*dt|du.*du|dv.*dv|dw.*dw|dx.*dx|dy.*dy|dz.*dz|ea.*ea|eb.*eb|ec.*ec|ed.*ed|ee.*ee|ef.*ef|eg.*eg|eh.*eh|ei.*ei|ej.*ej|ek.*ek|el.*el|em.*em|en.*en|eo.*eo|ep.*ep|eq.*eq|er.*er|es.*es|et.*et|eu.*eu|ev.*ev|ew.*ew|ex.*ex|ey.*ey|ez.*ez|fa.*fa|fb.*fb|fc.*fc|fd.*fd|fe.*fe|ff.*ff|fg.*fg|fh.*fh|fi.*fi|fj.*fj|fk.*fk|fl.*fl|fm.*fm|fn.*fn|fo.*fo|fp.*fp|fq.*fq|fr.*fr|fs.*fs|ft.*ft|fu.*fu|fv.*fv|fw.*fw|fx.*fx|fy.*fy|fz.*fz|ga.*ga|gb.*gb|gc.*gc|gd.*gd|ge.*ge|gf.*gf|gg.*gg|gh.*gh|gi.*gi|gj.*gj|gk.*gk|gl.*gl|gm.*gm|gn.*gn|go.*go|gp.*gp|gq.*gq|gr.*gr|gs.*gs|gt.*gt|gu.*gu|gv.*gv|gw.*gw|gx.*gx|gy.*gy|gz.*gz|ha.*ha|hb.*hb|hc.*hc|hd.*hd|he.*he|hf.*hf|hg.*hg|hh.*hh|hi.*hi|hj.*hj|hk.*hk|hl.*hl|hm.*hm|hn.*hn|ho.*ho|hp.*hp|hq.*hq|hr.*hr|hs.*hs|ht.*ht|hu.*hu|hv.*hv|hw.*hw|hx.*hx|hy.*hy|hz.*hz|ia.*ia|ib.*ib|ic.*ic|id.*id|ie.*ie|if.*if|ig.*ig|ih.*ih|ii.*ii|ij.*ij|ik.*ik|il.*il|im.*im|in.*in|io.*io|ip.*ip|iq.*iq|ir.*ir|is.*is|it.*it|iu.*iu|iv.*iv|iw.*iw|ix.*ix|iy.*iy|iz.*iz|ja.*ja|jb.*jb|jc.*jc|jd.*jd|je.*je|jf.*jf|jg.*jg|jh.*jh|ji.*ji|jj.*jj|jk.*jk|jl.*jl|jm.*jm|jn.*jn|jo.*jo|jp.*jp|jq.*jq|jr.*jr|js.*js|jt.*jt|ju.*ju|jv.*jv|jw.*jw|jx.*jx|jy.*jy|jz.*jz|ka.*ka|kb.*kb|kc.*kc|kd.*kd|ke.*ke|kf.*kf|kg.*kg|kh.*kh|ki.*ki|kj.*kj|kk.*kk|kl.*kl|km.*km|kn.*kn|ko.*ko|kp.*kp|kq.*kq|kr.*kr|ks.*ks|kt.*kt|ku.*ku|kv.*kv|kw.*kw|kx.*kx|ky.*ky|kz.*kz|la.*la|lb.*lb|lc.*lc|ld.*ld|le.*le|lf.*lf|lg.*lg|lh.*lh|li.*li|lj.*lj|lk.*lk|ll.*ll|lm.*lm|ln.*ln|lo.*lo|lp.*lp|lq.*lq|lr.*lr|ls.*ls|lt.*lt|lu.*lu|lv.*lv|lw.*lw|lx.*lx|ly.*ly|lz.*lz|ma.*ma|mb.*mb|mc.*mc|md.*md|me.*me|mf.*mf|mg.*mg|mh.*mh|mi.*mi|mj.*mj|mk.*mk|ml.*ml|mm.*mm|mn.*mn|mo.*mo|mp.*mp|mq.*mq|mr.*mr|ms.*ms|mt.*mt|mu.*mu|mv.*mv|mw.*mw|mx.*mx|my.*my|mz.*mz|na.*na|nb.*nb|nc.*nc|nd.*nd|ne.*ne|nf.*nf|ng.*ng|nh.*nh|ni.*ni|nj.*nj|nk.*nk|nl.*nl|nm.*nm|nn.*nn|no.*no|np.*np|nq.*nq|nr.*nr|ns.*ns|nt.*nt|nu.*nu|nv.*nv|nw.*nw|nx.*nx|ny.*ny|nz.*nz|oa.*oa|ob.*ob|oc.*oc|od.*od|oe.*oe|of.*of|og.*og|oh.*oh|oi.*oi|oj.*oj|ok.*ok|ol.*ol|om.*om|on.*on|oo.*oo|op.*op|oq.*oq|or.*or|os.*os|ot.*ot|ou.*ou|ov.*ov|ow.*ow|ox.*ox|oy.*oy|oz.*oz|pa.*pa|pb.*pb|pc.*pc|pd.*pd|pe.*pe|pf.*pf|pg.*pg|ph.*ph|pi.*pi|pj.*pj|pk.*pk|pl.*pl|pm.*pm|pn.*pn|po.*po|pp.*pp|pq.*pq|pr.*pr|ps.*ps|pt.*pt|pu.*pu|pv.*pv|pw.*pw|px.*px|py.*py|pz.*pz|qa.*qa|qb.*qb|qc.*qc|qd.*qd|qe.*qe|qf.*qf|qg.*qg|qh.*qh|qi.*qi|qj.*qj|qk.*qk|ql.*ql|qm.*qm|qn.*qn|qo.*qo|qp.*qp|qq.*qq|qr.*qr|qs.*qs|qt.*qt|qu.*qu|qv.*qv|qw.*qw|qx.*qx|qy.*qy|qz.*qz|ra.*ra|rb.*rb|rc.*rc|rd.*rd|re.*re|rf.*rf|rg.*rg|rh.*rh|ri.*ri|rj.*rj|rk.*rk|rl.*rl|rm.*rm|rn.*rn|ro.*ro|rp.*rp|rq.*rq|rr.*rr|rs.*rs|rt.*rt|ru.*ru|rv.*rv|rw.*rw|rx.*rx|ry.*ry|rz.*rz|sa.*sa|sb.*sb|sc.*sc|sd.*sd|se.*se|sf.*sf|sg.*sg|sh.*sh|si.*si|sj.*sj|sk.*sk|sl.*sl|sm.*sm|sn.*sn|so.*so|sp.*sp|sq.*sq|sr.*sr|ss.*ss|st.*st|su.*su|sv.*sv|sw.*sw|sx.*sx|sy.*sy|sz.*sz|ta.*ta|tb.*tb|tc.*tc|td.*td|te.*te|tf.*tf|tg.*tg|th.*th|ti.*ti|tj.*tj|tk.*tk|tl.*tl|tm.*tm|tn.*tn|to.*to|tp.*tp|tq.*tq|tr.*tr|ts.*ts|tt.*tt|tu.*tu|tv.*tv|tw.*tw|tx.*tx|ty.*ty|tz.*tz|ua.*ua|ub.*ub|uc.*uc|ud.*ud|ue.*ue|uf.*uf|ug.*ug|uh.*uh|ui.*ui|uj.*uj|uk.*uk|ul.*ul|um.*um|un.*un|uo.*uo|up.*up|uq.*uq|ur.*ur|us.*us|ut.*ut|uu.*uu|uv.*uv|uw.*uw|ux.*ux|uy.*uy|uz.*uz|va.*va|vb.*vb|vc.*vc|vd.*vd|ve.*ve|vf.*vf|vg.*vg|vh.*vh|vi.*vi|vj.*vj|vk.*vk|vl.*vl|vm.*vm|vn.*vn|vo.*vo|vp.*vp|vq.*vq|vr.*vr|vs.*vs|vt.*vt|vu.*vu|vv.*vv|vw.*vw|vx.*vx|vy.*vy|vz.*vz|wa.*wa|wb.*wb|wc.*wc|wd.*wd|we.*we|wf.*wf|wg.*wg|wh.*wh|wi.*wi|wj.*wj|wk.*wk|wl.*wl|wm.*wm|wn.*wn|wo.*wo|wp.*wp|wq.*wq|wr.*wr|ws.*ws|wt.*wt|wu.*wu|wv.*wv|ww.*ww|wx.*wx|wy.*wy|wz.*wz|xa.*xa|xb.*xb|xc.*xc|xd.*xd|xe.*xe|xf.*xf|xg.*xg|xh.*xh|xi.*xi|xj.*xj|xk.*xk|xl.*xl|xm.*xm|xn.*xn|xo.*xo|xp.*xp|xq.*xq|xr.*xr|xs.*xs|xt.*xt|xu.*xu|xv.*xv|xw.*xw|xx.*xx|xy.*xy|xz.*xz|ya.*ya|yb.*yb|yc.*yc|yd.*yd|ye.*ye|yf.*yf|yg.*yg|yh.*yh|yi.*yi|yj.*yj|yk.*yk|yl.*yl|ym.*ym|yn.*yn|yo.*yo|yp.*yp|yq.*yq|yr.*yr|ys.*ys|yt.*yt|yu.*yu|yv.*yv|yw.*yw|yx.*yx|yy.*yy|yz.*yz|za.*za|zb.*zb|zc.*zc|zd.*zd|ze.*ze|zf.*zf|zg.*zg|zh.*zh|zi.*zi|zj.*zj|zk.*zk|zl.*zl|zm.*zm|zn.*zn|zo.*zo|zp.*zp|zq.*zq|zr.*zr|zs.*zs|zt.*zt|zu.*zu|zv.*zv|zw.*zw|zx.*zx|zy.*zy|zz.*zy")
	// matches := twoPairs.FindAllString(line, -1)
	// return len(matches) >= 1

	indexMap := make(map[string]int)

	for index := 0; index < len(line)-1; index++ {

		pair := line[index : index+2]
		prevIndex, isPresent := indexMap[pair]
		if isPresent && prevIndex < index-1 {
			return true
		}

		if !isPresent {
			indexMap[pair] = index
		}
	}
	return false
}

func oneGap(line string) bool {
	// oneGap := regexp.MustCompile("a.a|b.b|c.c|d.d|e.e|f.f|g.g|h.h|i.i|j.j|k.k|l.l|m.m|n.n|o.o|p.p|q.q|r.r|s.s|t.t|u.u|v.v|w.w|x.x|y.y|z.z")
	// matches := oneGap.FindAllString(line, -1)
	// return len(matches) >= 1

	for index := 0; index < len(line)-2; index++ {
		if line[index] == line[index+2] {
			return true
		}
	}
	return false
}

func main() {
	util.Time(func() { solve(part1) }, "Part1")
	util.Time(func() { solve(part2) }, "Part2")
}

func test() {
	println(true == isNiceString("qjhvhtzxzqqjkmpb", true))
	println(true == isNiceString("xxyxx", true))
	println(false == isNiceString("uurcxstgmygtbstg", true))
	println(false == isNiceString("ieodomkazucvgmuy", true))

}
